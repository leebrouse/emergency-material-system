package model

import (
	"time"

	"gorm.io/gorm"
)

// Material 物资元数据
type Material struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"size:100;not null;index" json:"name"`
	Category    string         `gorm:"size:50;index" json:"category"`
	Specs       string         `gorm:"size:100" json:"specs"`          // 规格
	Unit        string         `gorm:"size:20" json:"unit"`            // 单位
	Quantity    int64          `gorm:"-" json:"quantity"`              // 聚合库存数量 (非数据库字段)
	MinStock    int64          `json:"min_stock"`                      // 安全库存
	BatchNumber string         `gorm:"size:50;index" json:"batch_num"` // 批次号
	ExpiryDate  *time.Time     `json:"expiry_date"`                    // 有效期
	Description string         `gorm:"size:255" json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName 指定物资表名
func (Material) TableName() string {
	return "materials"
}

// Inventory 库存实体
type Inventory struct {
	ID                  uint      `gorm:"primarykey" json:"id"`
	MaterialID          uint      `gorm:"uniqueIndex:idx_mat_loc" json:"material_id"`
	Material            Material  `gorm:"foreignKey:MaterialID" json:"material"`
	WarehouseLocation   string    `gorm:"size:100;uniqueIndex:idx_mat_loc" json:"location"` // 库位
	Quantity            int64     `gorm:"not null;default:0" json:"quantity"`               // 实际库存
	LockedQuantity      int64     `gorm:"not null;default:0" json:"locked_quantity"`        // 占用库存
	StockAlertThreshold int64     `gorm:"default:10" json:"alert_threshold"`                // 预警阈值
	Version             int64     `gorm:"default:0" json:"version"`                         // 乐观锁版本号
	UpdatedAt           time.Time `json:"updated_at"`
}

// TableName 指定库存表名
func (Inventory) TableName() string {
	return "inventory"
}

// StockLog 库存变更流水
type StockLog struct {
	ID             uint      `gorm:"primarykey" json:"id"`
	MaterialID     uint      `gorm:"index" json:"material_id"`
	InventoryID    uint      `gorm:"index" json:"inventory_id"`
	Type           LogType   `gorm:"size:20;index" json:"type"` // Inbound, Outbound, Transfer
	QuantityChange int64     `json:"quantity_change"`           // 变动数量
	BalanceAfter   int64     `json:"balance_after"`             // 变动后余额
	OperatorID     uint      `json:"operator_id"`               // 操作人ID
	RelatedOrderID string    `gorm:"size:50;index" json:"related_order_id"`
	Remark         string    `gorm:"size:255" json:"remark"`
	CreatedAt      time.Time `json:"created_at"`
}

// TableName 指定库存流水表名
func (StockLog) TableName() string {
	return "stock_logs"
}

type LogType string

const (
	LogTypeInbound  LogType = "Inbound"
	LogTypeOutbound LogType = "Outbound"
	LogTypeTransfer LogType = "Transfer"
)
