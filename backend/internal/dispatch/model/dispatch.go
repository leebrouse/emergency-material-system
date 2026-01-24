package model

import (
	"time"

	"gorm.io/gorm"
)

type UrgencyLevel string

const (
	UrgencyL1 UrgencyLevel = "L1"
	UrgencyL2 UrgencyLevel = "L2"
	UrgencyL3 UrgencyLevel = "L3"
)

type RequestStatus string

const (
	StatusPending  RequestStatus = "Pending"
	StatusAuditing RequestStatus = "Auditing"
	StatusApproved RequestStatus = "Approved"
	StatusRejected RequestStatus = "Rejected"
	StatusDispatch RequestStatus = "Dispatching"
	StatusShipping RequestStatus = "Shipping"
	StatusSigned   RequestStatus = "Signed"
)

// DemandRequest 需求申报单
type DemandRequest struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	MaterialID  uint           `gorm:"index" json:"material_id"`
	Quantity    int64          `json:"quantity"`
	Urgency     UrgencyLevel   `gorm:"size:10;index" json:"urgency"`
	TargetArea  string         `gorm:"size:100" json:"target_area"`
	Status      RequestStatus  `gorm:"size:20;index" json:"status"`
	Description string         `gorm:"size:255" json:"description"`
	AuditRemark string         `gorm:"size:255" json:"audit_remark"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// DispatchTask 配送任务单
type DispatchTask struct {
	ID          uint       `gorm:"primarykey" json:"id"`
	RequestID   uint       `gorm:"index" json:"request_id"`
	Status      TaskStatus `gorm:"size:20;index" json:"status"`
	Operator    string     `gorm:"size:50" json:"operator"`
	LogisticsID string     `gorm:"size:50" json:"logistics_id"` // 与物流系统关联
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type TaskStatus string

const (
	TaskStatusCreated   TaskStatus = "Created"
	TaskStatusInTransit TaskStatus = "InTransit"
	TaskStatusDelivered TaskStatus = "Delivered"
)

// DispatchLog 全链路追溯日志
type DispatchLog struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	RequestID  uint      `gorm:"index" json:"request_id"`
	PrevStatus string    `gorm:"size:20" json:"prev_status"`
	CurrStatus string    `gorm:"size:20" json:"curr_status"`
	OperatorID uint      `json:"operator_id"`
	Action     string    `gorm:"size:100" json:"action"`
	Remark     string    `gorm:"size:255" json:"remark"`
	CreatedAt  time.Time `json:"created_at"`
}

func (DemandRequest) TableName() string { return "demand_requests" }
func (DispatchTask) TableName() string  { return "dispatch_tasks" }
func (DispatchLog) TableName() string   { return "dispatch_logs" }
