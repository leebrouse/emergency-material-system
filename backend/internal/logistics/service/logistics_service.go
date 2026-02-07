package service

import (
	"context"
	"time"

	"github.com/emergency-material-system/backend/internal/common/utils"
	"github.com/emergency-material-system/backend/internal/logistics/model"
	"github.com/emergency-material-system/backend/internal/logistics/repository"
)

// LogisticsService 物流服务接口
type LogisticsService interface {
	GetTracking(ctx context.Context, id uint) (*model.Tracking, error)
	GetTrajectory(ctx context.Context, id uint) (*model.Tracking, error)
	CreateTracking(ctx context.Context, requestID uint, description, status string) (*model.Tracking, error)
	UpdateTracking(ctx context.Context, id uint, status, description string) error
	RecordTrajectoryNode(ctx context.Context, trackingID uint, location string, lat, lng float64, status, description string) error
}

// logisticsService 物流服务实现
type logisticsService struct {
	repo        repository.TrackingRepository
	gaodeClient *utils.GaodeMapsClient
}

// NewLogisticsService 创建物流服务
func NewLogisticsService(repo repository.TrackingRepository) LogisticsService {
	return &logisticsService{
		repo:        repo,
		gaodeClient: utils.NewGaodeMapsClient(),
	}
}

// GetTracking 获取物流追踪信息
func (s *logisticsService) GetTracking(ctx context.Context, id uint) (*model.Tracking, error) {
	return s.repo.GetByID(ctx, id)
}

// GetTrajectory 获取带完整轨迹的物流追踪信息 (可视化使用)
func (s *logisticsService) GetTrajectory(ctx context.Context, id uint) (*model.Tracking, error) {
	return s.repo.GetWithNodes(ctx, id)
}

// CreateTracking 创建物流追踪记录
func (s *logisticsService) CreateTracking(ctx context.Context, requestID uint, description, status string) (*model.Tracking, error) {
	tracking := &model.Tracking{
		RequestID:   requestID,
		Description: description,
		Status:      model.TrackingStatus(status),
		TrackedAt:   time.Now(),
	}

	if err := s.repo.Create(ctx, tracking); err != nil {
		return nil, err
	}

	return tracking, nil
}

// UpdateTracking 更新物流追踪状态
func (s *logisticsService) UpdateTracking(ctx context.Context, id uint, status, description string) error {
	tracking, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	tracking.Status = model.TrackingStatus(status)
	if description != "" {
		tracking.Description = description
	}

	return s.repo.Update(ctx, tracking)
}

// RecordTrajectoryNode 记录轨迹节点信息 (带百度地图 API 数据自动补全)
func (s *logisticsService) RecordTrajectoryNode(ctx context.Context, trackingID uint, location string, lat, lng float64, status, description string) error {
	// 数据补全逻辑
	if (lat == 0 && lng == 0) && location != "" {
		// 只有地址，没有经纬度 -> 地理编码
		if fetchedLat, fetchedLng, err := s.gaodeClient.GetCoordinatesFromAddress(location); err == nil {
			lat = fetchedLat
			lng = fetchedLng
		}
	} else if location == "" && (lat != 0 || lng != 0) {
		// 只有经纬度，没有地址 -> 逆地理编码
		if fetchedAddr, err := s.gaodeClient.GetAddressFromCoordinates(lat, lng); err == nil {
			location = fetchedAddr
		}
	}

	node := &model.TrackingNode{
		TrackingID:  trackingID,
		Location:    location,
		Latitude:    lat,
		Longitude:   lng,
		Status:      status,
		Description: description,
		TrackedAt:   time.Now(),
	}

	if err := s.repo.AddNode(ctx, node); err != nil {
		return err
	}

	// 同时更新主表的当前位置和状态
	tracking, err := s.repo.GetByID(ctx, trackingID)
	if err == nil {
		tracking.CurrentLoc = location
		if status != "" {
			tracking.Status = model.TrackingStatus(status)
		}
		_ = s.repo.Update(ctx, tracking)
	}

	return nil
}
