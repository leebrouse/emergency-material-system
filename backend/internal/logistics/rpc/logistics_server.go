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
	tracking, err := s.logisticsService.GetTrajectory(ctx, uint(req.OrderId))
	if err != nil {
		return nil, err
	}

	res := &logistics.GetTrackingResponse{
		Tracking: &logistics.Tracking{
			Id:         int64(tracking.ID),
			OrderId:    int64(tracking.RequestID),
			VehicleId:  "VEH001", // TODO: 从数据库中获取真实的车辆/司机信息
			DriverInfo: "司机张三",
			Status:     string(tracking.Status),
			CreatedAt:  tracking.CreatedAt.Unix(),
			UpdatedAt:  tracking.UpdatedAt.Unix(),
		},
	}

	if tracking.CurrentLoc != "" {
		res.Tracking.CurrentLocation = &logistics.TrackingPoint{
			Address:   tracking.CurrentLoc,
			Timestamp: tracking.TrackedAt.Unix(),
		}
	}

	return res, nil
}

// UpdateLocation 更新物流位置 (记录轨迹节点)
func (s *LogisticsRPCServer) UpdateLocation(ctx context.Context, req *logistics.UpdateLocationRequest) (*logistics.UpdateLocationResponse, error) {
	err := s.logisticsService.RecordTrajectoryNode(ctx, uint(req.TrackingId), req.Address, req.Latitude, req.Longitude, "in_transit", "位置更新")
	if err != nil {
		return nil, err
	}

	tracking, err := s.logisticsService.GetTracking(ctx, uint(req.TrackingId))
	if err != nil {
		return nil, err
	}

	return &logistics.UpdateLocationResponse{
		Tracking: &logistics.Tracking{
			Id:        int64(tracking.ID),
			OrderId:   int64(tracking.RequestID),
			Status:    string(tracking.Status),
			CreatedAt: tracking.CreatedAt.Unix(),
			UpdatedAt: tracking.UpdatedAt.Unix(),
			CurrentLocation: &logistics.TrackingPoint{
				Latitude:  req.Latitude,
				Longitude: req.Longitude,
				Address:   req.Address,
				Timestamp: time.Now().Unix(),
			},
		},
	}, nil
}

// GetTrackingHistory 获取物流历史轨迹点
func (s *LogisticsRPCServer) GetTrackingHistory(ctx context.Context, req *logistics.GetTrackingHistoryRequest) (*logistics.GetTrackingHistoryResponse, error) {
	tracking, err := s.logisticsService.GetTrajectory(ctx, uint(req.OrderId))
	if err != nil {
		return nil, err
	}

	points := make([]*logistics.TrackingPoint, 0, len(tracking.Nodes))
	for _, node := range tracking.Nodes {
		points = append(points, &logistics.TrackingPoint{
			Id:         int64(node.ID),
			TrackingId: int64(node.TrackingID),
			Latitude:   node.Latitude,
			Longitude:  node.Longitude,
			Address:    node.Location,
			Timestamp:  node.TrackedAt.Unix(),
		})
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
