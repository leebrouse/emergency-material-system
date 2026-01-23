package repository

import (
	"context"

	"github.com/emergency-material-system/backend/internal/stock/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StockRepository 综合库存仓库接口
type StockRepository interface {
	// Material operations
	CreateMaterial(ctx context.Context, m *model.Material) error
	GetMaterialByID(ctx context.Context, id uint) (*model.Material, error)
	ListMaterials(ctx context.Context, offset, limit int, query string) ([]*model.Material, int64, error)
	UpdateMaterial(ctx context.Context, m *model.Material) error

	// Inventory operations
	GetInventory(ctx context.Context, materialID uint, location string) (*model.Inventory, error)
	GetInventoryForUpdate(ctx context.Context, materialID uint, location string) (*model.Inventory, error)
	UpsertInventory(ctx context.Context, inv *model.Inventory) error
	ListInventory(ctx context.Context, offset, limit int) ([]*model.Inventory, int64, error)
	GetInventoryStats(ctx context.Context) ([]map[string]interface{}, error)

	// StockLog operations
	CreateStockLog(ctx context.Context, log *model.StockLog) error
	ListStockLogs(ctx context.Context, materialID uint, offset, limit int) ([]*model.StockLog, int64, error)

	// Transaction support
	Transaction(ctx context.Context, fn func(repo StockRepository) error) error
	WithTx(tx *gorm.DB) StockRepository
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockRepository{db: db}
}

func (r *stockRepository) WithTx(tx *gorm.DB) StockRepository {
	return &stockRepository{db: tx}
}

func (r *stockRepository) Transaction(ctx context.Context, fn func(repo StockRepository) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(r.WithTx(tx))
	})
}

// Material Impls
func (r *stockRepository) CreateMaterial(ctx context.Context, m *model.Material) error {
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *stockRepository) GetMaterialByID(ctx context.Context, id uint) (*model.Material, error) {
	var m model.Material
	err := r.db.WithContext(ctx).First(&m, id).Error
	return &m, err
}

func (r *stockRepository) ListMaterials(ctx context.Context, offset, limit int, search string) ([]*model.Material, int64, error) {
	var list []*model.Material
	var total int64
	db := r.db.WithContext(ctx).Model(&model.Material{})
	if search != "" {
		db = db.Where("name LIKE ? OR category LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	db.Count(&total)
	err := db.Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func (r *stockRepository) UpdateMaterial(ctx context.Context, m *model.Material) error {
	return r.db.WithContext(ctx).Save(m).Error
}

// Inventory Impls
func (r *stockRepository) GetInventory(ctx context.Context, materialID uint, location string) (*model.Inventory, error) {
	var inv model.Inventory
	err := r.db.WithContext(ctx).Where("material_id = ? AND warehouse_location = ?", materialID, location).First(&inv).Error
	return &inv, err
}

func (r *stockRepository) GetInventoryForUpdate(ctx context.Context, materialID uint, location string) (*model.Inventory, error) {
	var inv model.Inventory
	err := r.db.WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("material_id = ? AND warehouse_location = ?", materialID, location).First(&inv).Error
	return &inv, err
}

func (r *stockRepository) UpsertInventory(ctx context.Context, inv *model.Inventory) error {
	return r.db.WithContext(ctx).Save(inv).Error
}

func (r *stockRepository) ListInventory(ctx context.Context, offset, limit int) ([]*model.Inventory, int64, error) {
	var list []*model.Inventory
	var total int64
	db := r.db.WithContext(ctx).Model(&model.Inventory{}).Preload("Material")
	db.Count(&total)
	err := db.Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func (r *stockRepository) GetInventoryStats(ctx context.Context) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	err := r.db.WithContext(ctx).Model(&model.Inventory{}).
		Select("material_id, sum(quantity) as total_quantity, sum(locked_quantity) as total_locked").
		Group("material_id").Scan(&results).Error
	return results, err
}

// StockLog Impls
func (r *stockRepository) CreateStockLog(ctx context.Context, log *model.StockLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}

func (r *stockRepository) ListStockLogs(ctx context.Context, materialID uint, offset, limit int) ([]*model.StockLog, int64, error) {
	var list []*model.StockLog
	var total int64
	db := r.db.WithContext(ctx).Model(&model.StockLog{})
	if materialID > 0 {
		db = db.Where("material_id = ?", materialID)
	}
	db.Count(&total)
	err := db.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}
