package service

import (
	"context"
	"errors"
	"time"

	"github.com/emergency-material-system/backend/internal/logistics/model"
)

// LogisticsService 物流服务接口
type LogisticsService interface {
	GetTracking(ctx context.Context, id uint) (*model.Tracking, error)
	CreateTracking(ctx context.Context, requestID uint, description, status string) (*model.Tracking, error)
	UpdateTracking(ctx context.Context, id uint, status, description string) error
}

// logisticsService 物流服务实现
type logisticsService struct{}

// NewLogisticsService 创建物流服务
func NewLogisticsService() LogisticsService {
	return &logisticsService{}
}

// GetTracking 获取物流追踪信息
func (s *logisticsService) GetTracking(ctx context.Context, id uint) (*model.Tracking, error) {

	if id == 1 {
		return &model.Tracking{
			ID:          1,
			RequestID:   1,
			Description: "物资已发出",
			Status:      model.TrackingStatusInTransit,
			TrackedAt:   time.Now(),
		}, nil
	}

	return nil, errors.New("tracking not found")
}

// CreateTracking 创建物流追踪记录
func (s *logisticsService) CreateTracking(ctx context.Context, requestID uint, description, status string) (*model.Tracking, error) {

	tracking := &model.Tracking{
		ID:          2, // 模拟ID
		RequestID:   requestID,
		Description: description,
		Status:      model.TrackingStatus(status),
		TrackedAt:   time.Now(),
	}

	return tracking, nil
}

// UpdateTracking 更新物流追踪状态
func (s *logisticsService) UpdateTracking(ctx context.Context, id uint, status, description string) error {
	// 模拟更新
	return nil
}
