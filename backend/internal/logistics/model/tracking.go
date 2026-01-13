package model

import (
	"time"

	"gorm.io/gorm"
)

// Tracking 物流追踪模型
type Tracking struct {
	ID          uint            `gorm:"primarykey" json:"id"`
	RequestID   uint            `gorm:"not null" json:"request_id"`
	Description string          `json:"description"`
	Status      TrackingStatus  `gorm:"not null" json:"status"`
	TrackedAt   time.Time       `json:"tracked_at"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   gorm.DeletedAt  `gorm:"index" json:"-"`
}

// TrackingStatus 物流追踪状态
type TrackingStatus string

const (
	TrackingStatusCreated   TrackingStatus = "created"
	TrackingStatusPicked    TrackingStatus = "picked"
	TrackingStatusInTransit TrackingStatus = "in_transit"
	TrackingStatusDelivered TrackingStatus = "delivered"
	TrackingStatusCancelled TrackingStatus = "cancelled"
)

// TableName 指定表名
func (Tracking) TableName() string {
	return "tracking"
}
