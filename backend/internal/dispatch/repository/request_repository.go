package repository

import (
	"context"

	"github.com/emergency-material-system/backend/internal/dispatch/model"

	"gorm.io/gorm"
)

// RequestRepository 需求申报仓库接口
type RequestRepository interface {
	Create(ctx context.Context, request *model.Request) error
	GetByID(ctx context.Context, id uint) (*model.Request, error)
	List(ctx context.Context, offset, limit int) ([]*model.Request, int64, error)
	Update(ctx context.Context, request *model.Request) error
	Delete(ctx context.Context, id uint) error
}

// requestRepository 需求申报仓库实现
type requestRepository struct {
	db *gorm.DB
}

// NewRequestRepository 创建需求申报仓库
func NewRequestRepository(db *gorm.DB) RequestRepository {
	return &requestRepository{db: db}
}

// Create 创建需求申报
func (r *requestRepository) Create(ctx context.Context, request *model.Request) error {
	return r.db.WithContext(ctx).Create(request).Error
}

// GetByID 根据ID获取需求申报
func (r *requestRepository) GetByID(ctx context.Context, id uint) (*model.Request, error) {
	var request model.Request
	err := r.db.WithContext(ctx).First(&request, id).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}

// List 获取需求申报列表
func (r *requestRepository) List(ctx context.Context, offset, limit int) ([]*model.Request, int64, error) {
	var requests []*model.Request
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Request{})
	query.Count(&total)

	err := query.Offset(offset).Limit(limit).Find(&requests).Error
	return requests, total, err
}

// Update 更新需求申报
func (r *requestRepository) Update(ctx context.Context, request *model.Request) error {
	return r.db.WithContext(ctx).Save(request).Error
}

// Delete 删除需求申报
func (r *requestRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Request{}, id).Error
}
