package service

import (
	"context"
	"errors"

	"github.com/emergency-material-system/backend/internal/stock/model"
)

// StockService 物资库存服务接口
type StockService interface {
	ListMaterials(ctx context.Context, page, pageSize int) ([]*model.Material, int64, error)
	GetMaterial(ctx context.Context, id uint) (*model.Material, error)
	// CreateMaterial(ctx context.Context, req stock.PostStockMaterialsJSONBody) (*model.Material, error)
	GetInventory(ctx context.Context) ([]*model.Inventory, error)
	// UpdateInventory(ctx context.Context, req stock.PutStockInventoryJSONBody) error
}

// stockService 物资库存服务实现
type stockService struct{}

// NewStockService 创建物资库存服务
func NewStockService() StockService {
	return &stockService{}
}

// NewMockStockService 创建模拟物资库存服务
func NewMockStockService() StockService {
	return &stockService{}
}

// ListMaterials 获取物资列表
func (s *stockService) ListMaterials(ctx context.Context, page, pageSize int) ([]*model.Material, int64, error) {
	// 返回模拟数据
	materials := []*model.Material{
		{
			ID:          1,
			Name:        "口罩",
			Description: &[]string{"N95口罩"}[0],
			Category:    &[]string{"医疗物资"}[0],
			Unit:        &[]string{"个"}[0],
			Status:      model.MaterialStatusActive,
		},
		{
			ID:          2,
			Name:        "手套",
			Description: &[]string{"医用手套"}[0],
			Category:    &[]string{"医疗物资"}[0],
			Unit:        &[]string{"双"}[0],
			Status:      model.MaterialStatusActive,
		},
	}

	return materials, 2, nil
}

// GetMaterial 获取物资详情
func (s *stockService) GetMaterial(ctx context.Context, id uint) (*model.Material, error) {
	// 返回模拟数据
	if id == 1 {
		return &model.Material{
			ID:          1,
			Name:        "口罩",
			Description: &[]string{"N95口罩"}[0],
			Category:    &[]string{"医疗物资"}[0],
			Unit:        &[]string{"个"}[0],
			Status:      model.MaterialStatusActive,
		}, nil
	}

	return nil, errors.New("material not found")
}

// CreateMaterial 创建物资
// func (s *stockService) CreateMaterial(ctx context.Context, req stock.PostStockMaterialsJSONBody) (*model.Material, error) {

// 	if req.Name == nil || *req.Name == "" {
// 		return nil, errors.New("material name is required")
// 	}

// 	// 模拟创建
// 	material := &model.Material{
// 		ID:          3, // 模拟ID
// 		Name:        *req.Name,
// 		Description: req.Description,
// 		Category:    req.Category,
// 		Unit:        req.Unit,
// 		Status:      model.MaterialStatusActive,
// 	}

// 	return material, nil
// }

// GetInventory 获取库存信息
func (s *stockService) GetInventory(ctx context.Context) ([]*model.Inventory, error) {
	// 返回模拟数据
	inventory := []*model.Inventory{
		{
			ID:         1,
			MaterialID: 1,
			Quantity:   1000,
			Location:   &[]string{"仓库A"}[0],
			Status:     model.InvStatusAvailable,
		},
		{
			ID:         2,
			MaterialID: 2,
			Quantity:   500,
			Location:   &[]string{"仓库B"}[0],
			Status:     model.InvStatusAvailable,
		},
	}

	return inventory, nil
}

// // UpdateInventory 更新库存
// func (s *stockService) UpdateInventory(ctx context.Context, req stock.PutStockInventoryJSONBody) error {
// 	// 模拟更新
// 	return nil
// }
