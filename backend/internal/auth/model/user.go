package model

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Username  string         `gorm:"size:191;uniqueIndex;not null" json:"username"`
	Password  string         `gorm:"size:255;not null" json:"-"` // 密码不返回给前端
	Email     string         `gorm:"size:191;uniqueIndex" json:"email"`
	Phone     string         `gorm:"size:20" json:"phone"`
	Status    UserStatus     `gorm:"size:20;default:active" json:"status"`
	Roles     []Role         `gorm:"many2many:user_roles;" json:"roles"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// UserStatus 用户状态
type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
	UserStatusBanned   UserStatus = "banned"
)

// Role 角色模型
type Role struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Name        string    `gorm:"size:191;uniqueIndex;not null" json:"name"`
	Description string    `gorm:"size:255" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const (
	RoleAdmin   = "admin"
	RoleManager = "manager"
	RoleRescue  = "rescue"
)

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

func (Role) TableName() string {
	return "roles"
}
