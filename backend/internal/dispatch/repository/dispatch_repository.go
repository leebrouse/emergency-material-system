package repository

import (
	"context"

	"github.com/emergency-material-system/backend/internal/dispatch/model"

	"gorm.io/gorm"
)

type DispatchRepository interface {
	// DemandRequest
	CreateRequest(ctx context.Context, req *model.DemandRequest) error
	GetRequestByID(ctx context.Context, id uint) (*model.DemandRequest, error)
	UpdateRequest(ctx context.Context, req *model.DemandRequest) error
	ListRequests(ctx context.Context, offset, limit int, status string) ([]*model.DemandRequest, int64, error)

	// DispatchTask
	CreateTask(ctx context.Context, task *model.DispatchTask) error
	ListTasks(ctx context.Context, offset, limit int) ([]*model.DispatchTask, int64, error)

	// Logs
	CreateLog(ctx context.Context, log *model.DispatchLog) error

	// Transaction
	Transaction(ctx context.Context, fn func(repo DispatchRepository) error) error
	WithTx(tx *gorm.DB) DispatchRepository
}

type dispatchRepository struct {
	db *gorm.DB
}

func NewDispatchRepository(db *gorm.DB) DispatchRepository {
	return &dispatchRepository{db: db}
}

func (r *dispatchRepository) WithTx(tx *gorm.DB) DispatchRepository {
	return &dispatchRepository{db: tx}
}

func (r *dispatchRepository) Transaction(ctx context.Context, fn func(repo DispatchRepository) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(r.WithTx(tx))
	})
}

func (r *dispatchRepository) CreateRequest(ctx context.Context, req *model.DemandRequest) error {
	return r.db.WithContext(ctx).Create(req).Error
}

func (r *dispatchRepository) GetRequestByID(ctx context.Context, id uint) (*model.DemandRequest, error) {
	var req model.DemandRequest
	err := r.db.WithContext(ctx).First(&req, id).Error
	return &req, err
}

func (r *dispatchRepository) UpdateRequest(ctx context.Context, req *model.DemandRequest) error {
	return r.db.WithContext(ctx).Save(req).Error
}

func (r *dispatchRepository) ListRequests(ctx context.Context, offset, limit int, status string) ([]*model.DemandRequest, int64, error) {
	var list []*model.DemandRequest
	var total int64
	db := r.db.WithContext(ctx).Model(&model.DemandRequest{})
	if status != "" {
		db = db.Where("status = ?", status)
	}
	db.Count(&total)
	err := db.Offset(offset).Limit(limit).Order("id desc").Find(&list).Error
	return list, total, err
}

func (r *dispatchRepository) CreateTask(ctx context.Context, task *model.DispatchTask) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *dispatchRepository) ListTasks(ctx context.Context, offset, limit int) ([]*model.DispatchTask, int64, error) {
	var list []*model.DispatchTask
	var total int64
	db := r.db.WithContext(ctx).Model(&model.DispatchTask{})
	db.Count(&total)
	err := db.Offset(offset).Limit(limit).Find(&list).Error
	return list, total, err
}

func (r *dispatchRepository) CreateLog(ctx context.Context, log *model.DispatchLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}
