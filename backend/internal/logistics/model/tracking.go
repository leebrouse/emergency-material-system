package model

import (
	"time"

	"gorm.io/gorm"
)

// Tracking 物流追踪模型
type Tracking struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	RequestID   uint           `gorm:"not null" json:"request_id"`
	Description string         `json:"description"`
	Status      TrackingStatus `gorm:"not null" json:"status"`
	CurrentLoc  string         `json:"current_location"` // 当前位置描述
	TrackedAt   time.Time      `json:"tracked_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Nodes       []TrackingNode `gorm:"foreignKey:TrackingID" json:"nodes"` // 轨迹节点
}

// TrackingNode 物流轨迹节点模型
type TrackingNode struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	TrackingID  uint      `gorm:"index;not null" json:"tracking_id"`
	Location    string    `json:"location"`    // 位置名称
	Longitude   float64   `json:"longitude"`   // 经度 (高德/百度地图)
	Latitude    float64   `json:"latitude"`    // 纬度
	Status      string    `json:"status"`      // 该节点的物流状态说明
	Description string    `json:"description"` // 详细描述
	TrackedAt   time.Time `json:"tracked_at"`  // 记录时间
	CreatedAt   time.Time `json:"created_at"`
}

// TableName 指定 TrackingNode 表名
func (TrackingNode) TableName() string {
	return "tracking_nodes"
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
