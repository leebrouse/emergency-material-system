package service

import (
	"context"

	"github.com/emergency-material-system/backend/internal/dispatch/model"
)

// DispatchService 调度服务接口
type DispatchService interface {
	ListRequests(ctx context.Context, page, pageSize int) ([]*model.Request, int64, error)
	GetRequest(ctx context.Context, id uint) (*model.Request, error)
	CreateRequest(ctx context.Context, req interface{}) (*model.Request, error)
	UpdateRequestStatus(ctx context.Context, id uint, req interface{}) error
}

// dispatchService 调度服务实现
type dispatchService struct{}

// NewDispatchService 创建调度服务
func NewDispatchService() DispatchService {
	return &dispatchService{}
}

// ListRequests 获取需求申报列表
func (s *dispatchService) ListRequests(ctx context.Context, page, pageSize int) ([]*model.Request, int64, error) {
	return nil, 0, nil
}

// GetRequest 获取需求申报详情
func (s *dispatchService) GetRequest(ctx context.Context, id uint) (*model.Request, error) {
	return nil, nil
}

// CreateRequest 创建需求申报
func (s *dispatchService) CreateRequest(ctx context.Context, req interface{}) (*model.Request, error) {
	// 1.调 grpc stock service

	// 2. 满足则创建，不然erro

	// 模拟创建需求申报
	return nil, nil
}

// UpdateRequestStatus 更新需求申报状态
func (s *dispatchService) UpdateRequestStatus(ctx context.Context, id uint, req interface{}) error {
	// 模拟更新
	return nil
}
