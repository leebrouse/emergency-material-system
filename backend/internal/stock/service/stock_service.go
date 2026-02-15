package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/emergency-material-system/backend/internal/stock/model"
	"github.com/emergency-material-system/backend/internal/stock/repository"
)

// DTOs for stock operations
type InboundRequest struct {
	MaterialID uint   `json:"material_id" binding:"required"`
	Location   string `json:"location" binding:"required"`
	Quantity   int64  `json:"quantity" binding:"required,gt=0"`
	OperatorID uint   `json:"operator_id"`
	Remark     string `json:"remark"`
}

type OutboundRequest struct {
	MaterialID uint   `json:"material_id" binding:"required"`
	Location   string `json:"location" binding:"required"`
	Quantity   int64  `json:"quantity" binding:"required,gt=0"`
	OperatorID uint   `json:"operator_id"`
	Remark     string `json:"remark"`
}

type TransferRequest struct {
	MaterialID   uint   `json:"material_id" binding:"required"`
	FromLocation string `json:"from_location" binding:"required"`
	ToLocation   string `json:"to_location" binding:"required"`
	Quantity     int64  `json:"quantity" binding:"required,gt=0"`
	OperatorID   uint   `json:"operator_id"`
	Remark       string `json:"remark"`
}

// StockService 物资库存服务接口
type StockService interface {
	// Material Management
	CreateMaterial(ctx context.Context, m *model.Material) error
	ListMaterials(ctx context.Context, page, pageSize int, search string) ([]*model.Material, int64, error)
	GetMaterial(ctx context.Context, id uint) (*model.Material, error)

	// Stock Operations
	Inbound(ctx context.Context, req *InboundRequest) error
	Outbound(ctx context.Context, req *OutboundRequest) error
	Transfer(ctx context.Context, req *TransferRequest) error

	// Inventory & Stats
	ListInventory(ctx context.Context, page, pageSize int) ([]*model.Inventory, int64, error)
	ListInventoryByMaterial(ctx context.Context, materialID uint) ([]*model.Inventory, error)
	GetInventoryStats(ctx context.Context) ([]map[string]interface{}, error)
	LockStock(ctx context.Context, requestID uint, items map[uint]int64) error
}

type stockService struct {
	repo repository.StockRepository
}

func NewStockService(repo repository.StockRepository) StockService {
	return &stockService{repo: repo}
}

func (s *stockService) CreateMaterial(ctx context.Context, m *model.Material) error {
	return s.repo.CreateMaterial(ctx, m)
}

func (s *stockService) ListMaterials(ctx context.Context, page, pageSize int, search string) ([]*model.Material, int64, error) {
	offset := (page - 1) * pageSize
	return s.repo.ListMaterials(ctx, offset, pageSize, search)
}

func (s *stockService) GetMaterial(ctx context.Context, id uint) (*model.Material, error) {
	return s.repo.GetMaterialByID(ctx, id)
}

func (s *stockService) Inbound(ctx context.Context, req *InboundRequest) error {
	return s.repo.Transaction(ctx, func(tx repository.StockRepository) error {
		inv, err := tx.GetInventoryForUpdate(ctx, req.MaterialID, req.Location)
		if err != nil {
			// If not found, create new record
			inv = &model.Inventory{
				MaterialID:        req.MaterialID,
				WarehouseLocation: req.Location,
				Quantity:          0,
			}
		}

		inv.Quantity += req.Quantity
		if err := tx.UpsertInventory(ctx, inv); err != nil {
			return err
		}

		// Log entry
		return tx.CreateStockLog(ctx, &model.StockLog{
			MaterialID:     req.MaterialID,
			InventoryID:    inv.ID,
			Type:           model.LogTypeInbound,
			QuantityChange: req.Quantity,
			BalanceAfter:   inv.Quantity,
			OperatorID:     req.OperatorID,
			Remark:         req.Remark,
		})
	})
}

func (s *stockService) Outbound(ctx context.Context, req *OutboundRequest) error {
	return s.repo.Transaction(ctx, func(tx repository.StockRepository) error {
		inv, err := tx.GetInventoryForUpdate(ctx, req.MaterialID, req.Location)
		if err != nil {
			return fmt.Errorf("inventory not found: %w", err)
		}

		if inv.Quantity < req.Quantity {
			return errors.New("insufficient stock")
		}

		inv.Quantity -= req.Quantity
		if err := tx.UpsertInventory(ctx, inv); err != nil {
			return err
		}

		// Check for stock alert
		if inv.Quantity < inv.StockAlertThreshold {
			s.triggerStockAlert(inv)
		}

		return tx.CreateStockLog(ctx, &model.StockLog{
			MaterialID:     req.MaterialID,
			InventoryID:    inv.ID,
			Type:           model.LogTypeOutbound,
			QuantityChange: -req.Quantity,
			BalanceAfter:   inv.Quantity,
			OperatorID:     req.OperatorID,
			Remark:         req.Remark,
		})
	})
}

func (s *stockService) Transfer(ctx context.Context, req *TransferRequest) error {
	return s.repo.Transaction(ctx, func(tx repository.StockRepository) error {
		// 1. Outbound from source
		sourceInv, err := tx.GetInventoryForUpdate(ctx, req.MaterialID, req.FromLocation)
		if err != nil {
			return fmt.Errorf("source inventory not found: %w", err)
		}
		if sourceInv.Quantity < req.Quantity {
			return errors.New("insufficient stock in source location")
		}
		sourceInv.Quantity -= req.Quantity
		tx.UpsertInventory(ctx, sourceInv)

		// 2. Inbound to target
		targetInv, err := tx.GetInventoryForUpdate(ctx, req.MaterialID, req.ToLocation)
		if err != nil {
			targetInv = &model.Inventory{
				MaterialID:        req.MaterialID,
				WarehouseLocation: req.ToLocation,
				Quantity:          0,
			}
		}
		targetInv.Quantity += req.Quantity
		tx.UpsertInventory(ctx, targetInv)

		// 3. Logs
		tx.CreateStockLog(ctx, &model.StockLog{
			MaterialID:     req.MaterialID,
			InventoryID:    sourceInv.ID,
			Type:           model.LogTypeTransfer,
			QuantityChange: -req.Quantity,
			BalanceAfter:   sourceInv.Quantity,
			OperatorID:     req.OperatorID,
			Remark:         fmt.Sprintf("Transfer to %s: %s", req.ToLocation, req.Remark),
		})

		return tx.CreateStockLog(ctx, &model.StockLog{
			MaterialID:     req.MaterialID,
			InventoryID:    targetInv.ID,
			Type:           model.LogTypeTransfer,
			QuantityChange: req.Quantity,
			BalanceAfter:   targetInv.Quantity,
			OperatorID:     req.OperatorID,
			Remark:         fmt.Sprintf("Transfer from %s: %s", req.FromLocation, req.Remark),
		})
	})
}

func (s *stockService) ListInventory(ctx context.Context, page, pageSize int) ([]*model.Inventory, int64, error) {
	offset := (page - 1) * pageSize
	return s.repo.ListInventory(ctx, offset, pageSize)
}

func (s *stockService) ListInventoryByMaterial(ctx context.Context, materialID uint) ([]*model.Inventory, error) {
	return s.repo.ListInventoryByMaterial(ctx, materialID)
}

func (s *stockService) GetInventoryStats(ctx context.Context) ([]map[string]interface{}, error) {
	return s.repo.GetInventoryStats(ctx)
}

func (s *stockService) LockStock(ctx context.Context, requestID uint, items map[uint]int64) error {
	return s.repo.Transaction(ctx, func(tx repository.StockRepository) error {
		for invID, qty := range items {
			inv, err := tx.GetInventoryByIDForUpdate(ctx, invID)
			if err != nil {
				return err
			}
			if inv.Quantity < qty {
				return fmt.Errorf("insufficient stock for inventory %d", invID)
			}
			inv.Quantity -= qty
			inv.LockedQuantity += qty
			if err := tx.UpsertInventory(ctx, inv); err != nil {
				return err
			}

			// Log
			tx.CreateStockLog(ctx, &model.StockLog{
				MaterialID:     inv.MaterialID,
				InventoryID:    inv.ID,
				Type:           "Lock",
				QuantityChange: -qty,
				BalanceAfter:   inv.Quantity,
				Remark:         fmt.Sprintf("Lock for Request #%d", requestID),
			})
		}
		return nil
	})
}

// triggerStockAlert 库存预警 Hook
func (s *stockService) triggerStockAlert(inv *model.Inventory) {
	// 实际应用中可发送至消息队列 (MQ) 或通过钉钉/邮件通知
	fmt.Printf("ALERT: Stock for material %d in location %s is low: %d (Threshold: %d)\n",
		inv.MaterialID, inv.WarehouseLocation, inv.Quantity, inv.StockAlertThreshold)
}
