package service

import (
	"context"
	"fmt"
	"sort"

	"github.com/emergency-material-system/backend/internal/common/genproto/stock"
	"github.com/emergency-material-system/backend/internal/dispatch/model"
	"github.com/emergency-material-system/backend/internal/dispatch/repository"
)

type dispatchService struct {
	repo        repository.DispatchRepository
	stockClient stock.StockServiceClient
}

func NewDispatchService(repo repository.DispatchRepository, stockClient stock.StockServiceClient) DispatchService {
	return &dispatchService{
		repo:        repo,
		stockClient: stockClient,
	}
}

func (s *dispatchService) CreateDemandRequest(ctx context.Context, req *model.DemandRequest) error {
	req.Status = model.StatusPending
	err := s.repo.CreateRequest(ctx, req)
	if err == nil {
		s.log(ctx, req.ID, "", string(model.StatusPending), "CreateRequest", "User reported demand")
	}
	return err
}

func (s *dispatchService) GetDemandRequest(ctx context.Context, id uint) (*model.DemandRequest, error) {
	return s.repo.GetRequestByID(ctx, id)
}

func (s *dispatchService) ListDemandRequests(ctx context.Context, page, pageSize int, status string) ([]*model.DemandRequest, int64, error) {
	offset := (page - 1) * pageSize
	return s.repo.ListRequests(ctx, offset, pageSize, status)
}

func (s *dispatchService) AuditDemandRequest(ctx context.Context, id uint, action string, remark string) error {
	return s.repo.Transaction(ctx, func(tx repository.DispatchRepository) error {
		req, err := tx.GetRequestByID(ctx, id)
		if err != nil {
			return err
		}

		// State Machine Check
		if req.Status != model.StatusPending && req.Status != model.StatusAuditing {
			return fmt.Errorf("invalid status for audit: %s", req.Status)
		}

		prevStatus := string(req.Status)
		var nextStatus model.RequestStatus
		if action == "approve" {
			nextStatus = model.StatusApproved
		} else {
			nextStatus = model.StatusRejected
		}

		req.Status = nextStatus
		req.AuditRemark = remark
		if err := tx.UpdateRequest(ctx, req); err != nil {
			return err
		}

		return tx.CreateLog(ctx, &model.DispatchLog{
			RequestID:  id,
			PrevStatus: prevStatus,
			CurrStatus: string(nextStatus),
			Action:     "Audit",
			Remark:     remark,
		})
	})
}

func (s *dispatchService) SuggestAllocation(ctx context.Context, requestID uint) ([]AllocationSuggestion, error) {
	req, err := s.repo.GetRequestByID(ctx, requestID)
	if err != nil {
		return nil, err
	}

	// 1. Fetch available stocks from gRPC Stock Service
	resp, err := s.stockClient.ListInventoryItems(ctx, &stock.ListInventoryItemsRequest{
		MaterialId: int64(req.MaterialID),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stock: %w", err)
	}

	items := resp.Items
	// 2. Sorting Logic: Priority to almost expired (FEFO), then by location stability
	sort.Slice(items, func(i, j int) bool {
		expiryI := items[i].ExpiryDate
		expiryJ := items[j].ExpiryDate

		// If one has no expiry (0), treat it as far future (last priority)
		if expiryI == 0 {
			expiryI = 4102416000 // Year 2100+
		}
		if expiryJ == 0 {
			expiryJ = 4102416000
		}

		if expiryI != expiryJ {
			return expiryI < expiryJ
		}
		// Tie-breaker: stable location (alphabetical)
		return items[i].Location < items[j].Location
	})

	// 3. Match quantities
	var suggestions []AllocationSuggestion
	remaining := req.Quantity
	for _, item := range items {
		if remaining <= 0 {
			break
		}
		available := item.Quantity - item.LockedQuantity
		if available <= 0 {
			continue
		}

		take := min(available, remaining)

		suggestions = append(suggestions, AllocationSuggestion{
			InventoryID: uint(item.Id),
			Location:    item.Location,
			BatchNum:    item.BatchNum,
			Quantity:    take,
		})
		remaining -= take
	}

	if remaining > 0 {
		return suggestions, fmt.Errorf("insufficient stock: missing %d", remaining)
	}

	return suggestions, nil
}

func (s *dispatchService) CreateDispatchTask(ctx context.Context, requestID uint, allocations []AllocationSuggestion) (uint, error) {
	var taskID uint
	err := s.repo.Transaction(ctx, func(tx repository.DispatchRepository) error {
		req, err := tx.GetRequestByID(ctx, requestID)
		if err != nil {
			return err
		}

		if req.Status != model.StatusApproved {
			return fmt.Errorf("request must be approved before dispatch, current: %s", req.Status)
		}

		// 1. Lock Stock via gRPC (Stock Service)
		lockItems := make([]*stock.StockLockItem, len(allocations))
		stockLockMap := make(map[uint]int64)
		for i, a := range allocations {
			lockItems[i] = &stock.StockLockItem{
				InventoryId: int64(a.InventoryID),
				Quantity:    a.Quantity,
			}
			stockLockMap[a.InventoryID] = a.Quantity
		}

		lockResp, err := s.stockClient.LockStock(ctx, &stock.LockStockRequest{
			RequestId: int64(requestID),
			Items:     lockItems,
		})
		if err != nil {
			return fmt.Errorf("failed to call lock stock gRPC: %w", err)
		}
		if !lockResp.Success {
			return fmt.Errorf("stock lock failed: %s", lockResp.Message)
		}

		// 2. Create Dispatch Task
		task := &model.DispatchTask{
			RequestID: requestID,
			Status:    model.TaskStatusCreated,
		}
		if err := tx.CreateTask(ctx, task); err != nil {
			return err
		}
		taskID = task.ID

		// 3. Update Request Status
		prevStatus := string(req.Status)
		req.Status = model.StatusDispatch
		tx.UpdateRequest(ctx, req)

		// 4. Log
		tx.CreateLog(ctx, &model.DispatchLog{
			RequestID:  requestID,
			PrevStatus: prevStatus,
			CurrStatus: string(model.StatusDispatch),
			Action:     "CreateTask",
			Remark:     "Dispatch task generated and stock locked",
		})

		// 5. Notify Logistics (Mock with Channel/Event)
		go s.notifyLogistics(task)

		return nil
	})

	return taskID, err
}

func (s *dispatchService) ListDispatchTasks(ctx context.Context, page, pageSize int) ([]*model.DispatchTask, int64, error) {
	offset := (page - 1) * pageSize
	return s.repo.ListTasks(ctx, offset, pageSize)
}

func (s *dispatchService) log(ctx context.Context, reqID uint, prev, curr, action, remark string) {
	_ = s.repo.CreateLog(ctx, &model.DispatchLog{
		RequestID:  reqID,
		PrevStatus: prev,
		CurrStatus: curr,
		Action:     action,
		Remark:     remark,
	})
}

func (s *dispatchService) notifyLogistics(task *model.DispatchTask) {
	fmt.Printf("EVENT: Notifying logistics for Task #%d\n", task.ID)
	// In reality, send to RabbitMQ/Kafka or call Logistics gRPC
}
