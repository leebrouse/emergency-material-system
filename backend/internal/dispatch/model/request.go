package model

import (
	"time"

	"gorm.io/gorm"
)

// Request 需求申报模型
type Request struct {
	ID           uint            `gorm:"primarykey" json:"id"`
	RequesterID  *uint           `json:"requester_id"`
	MaterialID   *uint           `json:"material_id"`
	Quantity     *int            `gorm:"not null" json:"quantity"`
	UrgencyLevel *string         `json:"urgency_level"`
	Description  *string         `json:"description"`
	Status       RequestStatus   `gorm:"default:pending" json:"status"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    gorm.DeletedAt  `gorm:"index" json:"-"`
}

// RequestStatus 需求申报状态
type RequestStatus string

const (
	RequestStatusPending   RequestStatus = "pending"
	RequestStatusApproved  RequestStatus = "approved"
	RequestStatusRejected  RequestStatus = "rejected"
	RequestStatusCompleted RequestStatus = "completed"
)

// TableName 指定表名
func (Request) TableName() string {
	return "requests"
}
