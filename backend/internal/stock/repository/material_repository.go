package repository

import (
	"context"

	"github.com/emergency-material-system/backend/internal/stock/model"

	"gorm.io/gorm"
)

// MaterialRepository 物资仓库接口
type MaterialRepository interface {
	Create(ctx context.Context, material *model.Material) error
	GetByID(ctx context.Context, id uint) (*model.Material, error)
	List(ctx context.Context, offset, limit int) ([]*model.Material, int64, error)
	Update(ctx context.Context, material *model.Material) error
	Delete(ctx context.Context, id uint) error
}

// InventoryRepository 库存仓库接口
type InventoryRepository interface {
	Create(ctx context.Context, inventory *model.Inventory) error
	GetByID(ctx context.Context, id uint) (*model.Inventory, error)
	List(ctx context.Context) ([]*model.Inventory, error)
	Update(ctx context.Context, inventory *model.Inventory) error
}

// materialRepository 物资仓库实现
type materialRepository struct {
	db *gorm.DB
}

// NewMaterialRepository 创建物资仓库
func NewMaterialRepository(db *gorm.DB) MaterialRepository {
	return &materialRepository{db: db}
}

// Create 创建物资
func (r *materialRepository) Create(ctx context.Context, material *model.Material) error {
	return r.db.WithContext(ctx).Create(material).Error
}

// GetByID 根据ID获取物资
func (r *materialRepository) GetByID(ctx context.Context, id uint) (*model.Material, error) {
	var material model.Material
	err := r.db.WithContext(ctx).First(&material, id).Error
	if err != nil {
		return nil, err
	}
	return &material, nil
}

// List 获取物资列表
func (r *materialRepository) List(ctx context.Context, offset, limit int) ([]*model.Material, int64, error) {
	var materials []*model.Material
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Material{})
	query.Count(&total)

	err := query.Offset(offset).Limit(limit).Find(&materials).Error
	return materials, total, err
}

// Update 更新物资
func (r *materialRepository) Update(ctx context.Context, material *model.Material) error {
	return r.db.WithContext(ctx).Save(material).Error
}

// Delete 删除物资
func (r *materialRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Material{}, id).Error
}

// inventoryRepository 库存仓库实现
type inventoryRepository struct {
	db *gorm.DB
}

// NewInventoryRepository 创建库存仓库
func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db: db}
}

// Create 创建库存记录
func (r *inventoryRepository) Create(ctx context.Context, inventory *model.Inventory) error {
	return r.db.WithContext(ctx).Create(inventory).Error
}

// GetByID 根据ID获取库存
func (r *inventoryRepository) GetByID(ctx context.Context, id uint) (*model.Inventory, error) {
	var inventory model.Inventory
	err := r.db.WithContext(ctx).Preload("Material").First(&inventory, id).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}

// List 获取库存列表
func (r *inventoryRepository) List(ctx context.Context) ([]*model.Inventory, error) {
	var inventories []*model.Inventory
	err := r.db.WithContext(ctx).Preload("Material").Find(&inventories).Error
	return inventories, err
}

// Update 更新库存
func (r *inventoryRepository) Update(ctx context.Context, inventory *model.Inventory) error {
	return r.db.WithContext(ctx).Save(inventory).Error
}
