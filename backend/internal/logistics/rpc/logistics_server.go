package rpc

import (
	"context"
	"time"

	"github.com/emergency-material-system/backend/internal/common/genproto/logistics"
	"github.com/emergency-material-system/backend/internal/logistics/service"

	"google.golang.org/grpc"
)

// LogisticsRPCServer 物流gRPC服务器
type LogisticsRPCServer struct {
	logistics.UnimplementedLogisticsServiceServer
	logisticsService service.LogisticsService
}

// NewLogisticsRPCServer 创建物流gRPC服务器
func NewLogisticsRPCServer(logisticsService service.LogisticsService) *LogisticsRPCServer {
	return &LogisticsRPCServer{
		logisticsService: logisticsService,
	}
}

// Register 注册gRPC服务
func (s *LogisticsRPCServer) Register(server *grpc.Server) {
	logistics.RegisterLogisticsServiceServer(server, s)
}

// GetTracking 获取物流轨迹
func (s *LogisticsRPCServer) GetTracking(ctx context.Context, req *logistics.GetTrackingRequest) (*logistics.GetTrackingResponse, error) {
	tracking, err := s.logisticsService.GetTracking(ctx, uint(req.OrderId))
	if err != nil {
		return nil, err
	}

	return &logistics.GetTrackingResponse{
		Tracking: &logistics.Tracking{
			Id:         int64(tracking.ID),
			OrderId:    int64(tracking.RequestID),
			VehicleId:  "VEH001", // 模拟车辆ID
			DriverInfo: "司机张三",   // 模拟司机信息
			Status:     string(tracking.Status),
			CurrentLocation: &logistics.TrackingPoint{
				Id:         int64(tracking.ID),
				TrackingId: int64(tracking.ID),
				Latitude:   39.9042, // 模拟位置
				Longitude:  116.4074,
				Address:    "北京市朝阳区",
				Timestamp:  tracking.TrackedAt.Unix(),
			},
			CreatedAt: tracking.CreatedAt.Unix(),
			UpdatedAt: tracking.UpdatedAt.Unix(),
		},
	}, nil
}

// UpdateLocation 更新物流位置
func (s *LogisticsRPCServer) UpdateLocation(ctx context.Context, req *logistics.UpdateLocationRequest) (*logistics.UpdateLocationResponse, error) {
	// 模拟更新位置逻辑
	// 这里可以调用实际的物流服务来更新位置信息

	return &logistics.UpdateLocationResponse{
		Tracking: &logistics.Tracking{
			Id:         req.TrackingId,
			OrderId:    1, // 模拟订单ID
			VehicleId:  "VEH001",
			DriverInfo: "司机张三",
			Status:     "in_transit",
			CurrentLocation: &logistics.TrackingPoint{
				Id:         req.TrackingId,
				TrackingId: req.TrackingId,
				Latitude:   req.Latitude,
				Longitude:  req.Longitude,
				Address:    req.Address,
				Timestamp:  time.Now().Unix(),
			},
			CreatedAt: time.Now().Unix(),
			UpdatedAt: time.Now().Unix(),
		},
	}, nil
}

// GetTrackingHistory 获取物流历史
func (s *LogisticsRPCServer) GetTrackingHistory(ctx context.Context, req *logistics.GetTrackingHistoryRequest) (*logistics.GetTrackingHistoryResponse, error) {
	// 模拟物流历史数据 - 根据订单ID返回轨迹历史
	points := []*logistics.TrackingPoint{
		{
			Id:         1,
			TrackingId: req.OrderId,
			Latitude:   39.9042,
			Longitude:  116.4074,
			Address:    "北京市朝阳区仓库",
			Timestamp:  time.Now().Add(-2 * time.Hour).Unix(),
		},
		{
			Id:         2,
			TrackingId: req.OrderId,
			Latitude:   39.9142,
			Longitude:  116.4174,
			Address:    "北京市海淀区中转站",
			Timestamp:  time.Now().Add(-1 * time.Hour).Unix(),
		},
		{
			Id:         3,
			TrackingId: req.OrderId,
			Latitude:   39.9242,
			Longitude:  116.4274,
			Address:    "北京市西城区配送点",
			Timestamp:  time.Now().Add(-30 * time.Minute).Unix(),
		},
	}

	return &logistics.GetTrackingHistoryResponse{
		Points: points,
	}, nil
}

// CreateTracking 创建物流记录
func (s *LogisticsRPCServer) CreateTracking(ctx context.Context, req *logistics.CreateTrackingRequest) (*logistics.CreateTrackingResponse, error) {
	// 使用现有的服务方法创建物流记录
	tracking, err := s.logisticsService.CreateTracking(ctx, uint(req.OrderId), "物流记录创建", "created")
	if err != nil {
		return nil, err
	}

	return &logistics.CreateTrackingResponse{
		Tracking: &logistics.Tracking{
			Id:         int64(tracking.ID),
			OrderId:    int64(tracking.RequestID),
			VehicleId:  req.VehicleId,
			DriverInfo: req.DriverInfo,
			Status:     string(tracking.Status),
			CreatedAt:  tracking.CreatedAt.Unix(),
			UpdatedAt:  tracking.UpdatedAt.Unix(),
		},
	}, nil
}
