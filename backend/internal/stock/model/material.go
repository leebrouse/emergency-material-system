package model

import (
	"time"

	"gorm.io/gorm"
)

// Material 物资模型
type Material struct {
	ID          uint            `gorm:"primarykey" json:"id"`
	Name        string          `gorm:"not null" json:"name"`
	Description *string         `json:"description"`
	Category    *string         `json:"category"`
	Unit        *string         `json:"unit"`
	Status      MaterialStatus  `gorm:"default:active" json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   gorm.DeletedAt  `gorm:"index" json:"-"`
}

// MaterialStatus 物资状态
type MaterialStatus string

const (
	MaterialStatusActive   MaterialStatus = "active"
	MaterialStatusInactive MaterialStatus = "inactive"
	MaterialStatusDeleted  MaterialStatus = "deleted"
)

// Inventory 库存模型
type Inventory struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	MaterialID uint      `gorm:"not null" json:"material_id"`
	Material   Material  `gorm:"foreignKey:MaterialID" json:"material"`
	Quantity   int       `gorm:"not null" json:"quantity"`
	Location   *string   `json:"location"`
	Status     InvStatus `gorm:"default:available" json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// InvStatus 库存状态
type InvStatus string

const (
	InvStatusAvailable InvStatus = "available"
	InvStatusReserved  InvStatus = "reserved"
	InvStatusOut       InvStatus = "out"
)

// TableName 指定表名
func (Material) TableName() string {
	return "materials"
}

func (Inventory) TableName() string {
	return "inventory"
}
