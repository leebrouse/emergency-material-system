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

// Implement mock/stub for GRPC methods to avoid compile errors
func (s *DispatchRPCServer) ListDemands(ctx context.Context, req *dispatch.ListDemandsRequest) (*dispatch.ListDemandsResponse, error) {
	return nil, fmt.Errorf("not implemented via gRPC")
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
