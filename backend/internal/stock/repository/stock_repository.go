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
	DeleteMaterial(ctx context.Context, id uint) error

	// Inventory operations
	GetInventory(ctx context.Context, materialID uint, location string) (*model.Inventory, error)
	GetInventoryForUpdate(ctx context.Context, materialID uint, location string) (*model.Inventory, error)
	UpsertInventory(ctx context.Context, inv *model.Inventory) error
	ListInventory(ctx context.Context, offset, limit int) ([]*model.Inventory, int64, error)
	GetInventoryStats(ctx context.Context) ([]map[string]interface{}, error)

	// StockLog operations
	CreateStockLog(ctx context.Context, log *model.StockLog) error
	ListStockLogs(ctx context.Context, materialID uint, logType string, offset, limit int) ([]*model.StockLog, int64, error)

	// Dispatch Support
	ListInventoryByMaterial(ctx context.Context, materialID uint) ([]*model.Inventory, error)
	GetInventoryByIDForUpdate(ctx context.Context, id uint) (*model.Inventory, error)

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
	if err != nil {
		return nil, 0, err
	}

	if len(list) > 0 {
		var ids []uint
		for _, m := range list {
			ids = append(ids, m.ID)
		}

		type Result struct {
			MaterialID uint
			Total      int64
		}
		var results []Result

		r.db.WithContext(ctx).Table("inventory").
			Select("material_id, SUM(quantity) as total").
			Where("material_id IN ?", ids).
			Group("material_id").
			Scan(&results)

		qtyMap := make(map[uint]int64)
		for _, res := range results {
			qtyMap[res.MaterialID] = res.Total
		}

		for i, m := range list {
			list[i].Quantity = qtyMap[m.ID]
		}
	}

	return list, total, nil
}

func (r *stockRepository) UpdateMaterial(ctx context.Context, m *model.Material) error {
	return r.db.WithContext(ctx).Save(m).Error
}

func (r *stockRepository) DeleteMaterial(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Material{}, id).Error
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

func (r *stockRepository) ListStockLogs(ctx context.Context, materialID uint, logType string, offset, limit int) ([]*model.StockLog, int64, error) {
	var list []*model.StockLog
	var total int64
	db := r.db.WithContext(ctx).Model(&model.StockLog{})
	if materialID > 0 {
		db = db.Where("material_id = ?", materialID)
	}
	if logType != "" {
		db = db.Where("type = ?", logType)
	}
	db.Count(&total)
	err := db.Order("id desc").Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func (r *stockRepository) ListInventoryByMaterial(ctx context.Context, materialID uint) ([]*model.Inventory, error) {
	var list []*model.Inventory
	err := r.db.WithContext(ctx).Preload("Material").Where("material_id = ?", materialID).Order("updated_at asc").Find(&list).Error
	return list, err
}

func (r *stockRepository) GetInventoryByIDForUpdate(ctx context.Context, id uint) (*model.Inventory, error) {
	var inv model.Inventory
	err := r.db.WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).First(&inv, id).Error
	return &inv, err
}
