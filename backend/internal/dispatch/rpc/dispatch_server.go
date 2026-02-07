package rpc

import (
	"context"
	"fmt"

	"github.com/emergency-material-system/backend/internal/common/genproto/dispatch"
	"github.com/emergency-material-system/backend/internal/dispatch/service"

	"google.golang.org/grpc"
)

// DispatchRPCServer 调度gRPC服务器
type DispatchRPCServer struct {
	dispatch.UnimplementedDispatchServiceServer
	dispatchService service.DispatchService
}

// NewDispatchRPCServer 创建调度gRPC服务器
func NewDispatchRPCServer(dispatchService service.DispatchService) *DispatchRPCServer {
	return &DispatchRPCServer{
		dispatchService: dispatchService,
	}
}

// Register 注册gRPC服务
func (s *DispatchRPCServer) Register(server *grpc.Server) {
	dispatch.RegisterDispatchServiceServer(server, s)
}

// ListDemands 获取所有需求列表 (用于统计)
func (s *DispatchRPCServer) ListDemands(ctx context.Context, req *dispatch.ListDemandsRequest) (*dispatch.ListDemandsResponse, error) {
	demands, _, err := s.dispatchService.ListDemandRequests(ctx, 1, 1000, "") // 获取前1000条记录
	if err != nil {
		return nil, err
	}

	var demandProtos []*dispatch.Demand
	for _, d := range demands {
		demandProtos = append(demandProtos, &dispatch.Demand{
			Id:          int64(d.ID),
			Location:    d.TargetArea,
			Priority:    string(d.Urgency),
			Status:      string(d.Status),
			Description: d.Description,
			CreatedAt:   d.CreatedAt.Unix(),
			UpdatedAt:   d.UpdatedAt.Unix(),
			Items: []*dispatch.DemandItem{
				{
					MaterialId: int64(d.MaterialID),
					Quantity:   d.Quantity,
				},
			},
		})
	}

	return &dispatch.ListDemandsResponse{
		Demands: demandProtos,
		Total:   int32(len(demands)),
	}, nil
}

func (s *DispatchRPCServer) CreateDemand(ctx context.Context, req *dispatch.CreateDemandRequest) (*dispatch.CreateDemandResponse, error) {
	return nil, fmt.Errorf("not implemented via gRPC")
}

func (s *DispatchRPCServer) GetDemand(ctx context.Context, req *dispatch.GetDemandRequest) (*dispatch.GetDemandResponse, error) {
	return nil, fmt.Errorf("not implemented via gRPC")
}

func (s *DispatchRPCServer) UpdateDemandStatus(ctx context.Context, req *dispatch.UpdateDemandStatusRequest) (*dispatch.UpdateDemandStatusResponse, error) {
	return nil, fmt.Errorf("not implemented via gRPC")
}

func (s *DispatchRPCServer) CreateDispatchOrder(ctx context.Context, req *dispatch.CreateDispatchOrderRequest) (*dispatch.CreateDispatchOrderResponse, error) {
	return nil, fmt.Errorf("not implemented via gRPC")
}

func (s *DispatchRPCServer) ListDispatchOrders(ctx context.Context, req *dispatch.ListDispatchOrdersRequest) (*dispatch.ListDispatchOrdersResponse, error) {
	return nil, fmt.Errorf("not implemented via gRPC")
}
