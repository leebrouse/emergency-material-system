package service

import (
	"context"
	"errors"

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

// NewMockDispatchService 创建模拟调度服务
func NewMockDispatchService() DispatchService {
	return &dispatchService{}
}

// ListRequests 获取需求申报列表
func (s *dispatchService) ListRequests(ctx context.Context, page, pageSize int) ([]*model.Request, int64, error) {
	// 返回模拟数据
	requests := []*model.Request{
		{
			ID:           1,
			RequesterID:  &[]uint{1}[0],
			MaterialID:   &[]uint{1}[0],
			Quantity:     &[]int{100}[0],
			UrgencyLevel: &[]string{"high"}[0],
			Description:  &[]string{"急需口罩"}[0],
			Status:       model.RequestStatusPending,
		},
	}

	return requests, 1, nil
}

// GetRequest 获取需求申报详情
func (s *dispatchService) GetRequest(ctx context.Context, id uint) (*model.Request, error) {
	if id == 1 {
		return &model.Request{
			ID:           1,
			RequesterID:  &[]uint{1}[0],
			MaterialID:   &[]uint{1}[0],
			Quantity:     &[]int{100}[0],
			UrgencyLevel: &[]string{"high"}[0],
			Description:  &[]string{"急需口罩"}[0],
			Status:       model.RequestStatusPending,
		}, nil
	}

	return nil, errors.New("request not found")
}

// CreateRequest 创建需求申报
func (s *dispatchService) CreateRequest(ctx context.Context, req interface{}) (*model.Request, error) {
	// 模拟创建需求申报
	request := &model.Request{
		ID:           2, // 模拟ID
		RequesterID:  &[]uint{1}[0],
		MaterialID:   &[]uint{1}[0],
		Quantity:     &[]int{100}[0],
		UrgencyLevel: &[]string{"high"}[0],
		Description:  &[]string{"模拟需求"}[0],
		Status:       model.RequestStatusPending,
	}
	return request, nil
}

// UpdateRequestStatus 更新需求申报状态
func (s *dispatchService) UpdateRequestStatus(ctx context.Context, id uint, req interface{}) error {
	// 模拟更新
	return nil
}
