package service

import (
	"context"

	"github.com/emergency-material-system/backend/internal/dispatch/model"
)

type AllocationSuggestion struct {
	InventoryID uint   `json:"inventory_id"`
	Location    string `json:"location"`
	BatchNum    string `json:"batch_num"`
	Quantity    int64  `json:"quantity"`
}

type DispatchService interface {
	// Demand Request Management
	CreateDemandRequest(ctx context.Context, req *model.DemandRequest) error
	GetDemandRequest(ctx context.Context, id uint) (*model.DemandRequest, error)
	ListDemandRequests(ctx context.Context, page, pageSize int, status string) ([]*model.DemandRequest, int64, error)

	// Audit & Workflow
	AuditDemandRequest(ctx context.Context, id uint, action string, remark string) error

	// Intelligent Allocation
	SuggestAllocation(ctx context.Context, requestID uint) ([]AllocationSuggestion, error)

	// Dispatch Execution
	CreateDispatchTask(ctx context.Context, requestID uint, allocations []AllocationSuggestion) (uint, error)
	ListDispatchTasks(ctx context.Context, page, pageSize int) ([]*model.DispatchTask, int64, error)
}
