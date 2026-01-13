package repository

import (
	"context"

	"github.com/emergency-material-system/backend/internal/logistics/model"

	"gorm.io/gorm"
)

// TrackingRepository 物流追踪仓库接口
type TrackingRepository interface {
	Create(ctx context.Context, tracking *model.Tracking) error
	GetByID(ctx context.Context, id uint) (*model.Tracking, error)
	GetByRequestID(ctx context.Context, requestID uint) ([]*model.Tracking, error)
	Update(ctx context.Context, tracking *model.Tracking) error
	Delete(ctx context.Context, id uint) error
}

// trackingRepository 物流追踪仓库实现
type trackingRepository struct {
	db *gorm.DB
}

// NewTrackingRepository 创建物流追踪仓库
func NewTrackingRepository(db *gorm.DB) TrackingRepository {
	return &trackingRepository{db: db}
}

// Create 创建物流追踪记录
func (r *trackingRepository) Create(ctx context.Context, tracking *model.Tracking) error {
	return r.db.WithContext(ctx).Create(tracking).Error
}

// GetByID 根据ID获取物流追踪
func (r *trackingRepository) GetByID(ctx context.Context, id uint) (*model.Tracking, error) {
	var tracking model.Tracking
	err := r.db.WithContext(ctx).First(&tracking, id).Error
	if err != nil {
		return nil, err
	}
	return &tracking, nil
}

// GetByRequestID 根据需求ID获取物流追踪记录
func (r *trackingRepository) GetByRequestID(ctx context.Context, requestID uint) ([]*model.Tracking, error) {
	var trackings []*model.Tracking
	err := r.db.WithContext(ctx).Where("request_id = ?", requestID).Order("tracked_at DESC").Find(&trackings).Error
	return trackings, err
}

// Update 更新物流追踪
func (r *trackingRepository) Update(ctx context.Context, tracking *model.Tracking) error {
	return r.db.WithContext(ctx).Save(tracking).Error
}

// Delete 删除物流追踪
func (r *trackingRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Tracking{}, id).Error
}
