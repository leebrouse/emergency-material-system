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

// ListDemands 获取需求列表
func (s *DispatchRPCServer) ListDemands(ctx context.Context, req *dispatch.ListDemandsRequest) (*dispatch.ListDemandsResponse, error) {

	requests, total, err := s.dispatchService.ListRequests(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}

	// 转换数据格式
	var requestProtos []*dispatch.Demand
	for _, r := range requests {
		requestProtos = append(requestProtos, &dispatch.Demand{
			Id:          int64(r.ID),
			Location:    "default",
			Priority:    *r.UrgencyLevel,
			Description: *r.Description,
			Status:      string(r.Status),
		})
	}

	return &dispatch.ListDemandsResponse{
		Demands: requestProtos,
		Total:   int32(total),
	}, nil
}

// CreateDemand 创建需求
func (s *DispatchRPCServer) CreateDemand(ctx context.Context, req *dispatch.CreateDemandRequest) (*dispatch.CreateDemandResponse, error) {

	// 转换请求格式 - 使用protobuf字段
	createReq := map[string]interface{}{
		"location":    req.Location,
		"priority":    req.Priority,
		"description": req.Description,
	}

	request, err := s.dispatchService.CreateRequest(ctx, createReq)
	if err != nil {
		return nil, err
	}

	return &dispatch.CreateDemandResponse{
		Demand: &dispatch.Demand{
			Id:          int64(request.ID),
			Location:    "default",
			Priority:    *request.UrgencyLevel,
			Description: *request.Description,
			Status:      string(request.Status),
		},
	}, nil
}

// GetDemand 获取需求详情
func (s *DispatchRPCServer) GetDemand(ctx context.Context, req *dispatch.GetDemandRequest) (*dispatch.GetDemandResponse, error) {

	request, err := s.dispatchService.GetRequest(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &dispatch.GetDemandResponse{
		Demand: &dispatch.Demand{
			Id:          int64(request.ID),
			Location:    "default",
			Priority:    *request.UrgencyLevel,
			Description: *request.Description,
			Status:      string(request.Status),
		},
	}, nil
}

// UpdateDemandStatus 更新需求状态
func (s *DispatchRPCServer) UpdateDemandStatus(ctx context.Context, req *dispatch.UpdateDemandStatusRequest) (*dispatch.UpdateDemandStatusResponse, error) {

	// 转换请求格式 - 使用简单的类型转换
	updateReq := map[string]interface{}{
		"status": req.Status,
	}

	err := s.dispatchService.UpdateRequestStatus(ctx, uint(req.Id), updateReq)
	if err != nil {
		return nil, err
	}

	return &dispatch.UpdateDemandStatusResponse{Demand: nil}, nil
}

// CreateDispatchOrder 创建调度订单
func (s *DispatchRPCServer) CreateDispatchOrder(ctx context.Context, req *dispatch.CreateDispatchOrderRequest) (*dispatch.CreateDispatchOrderResponse, error) {
	// 暂时返回未实现
	return nil, fmt.Errorf("not implemented")
}

// ListDispatchOrders 获取调度订单列表
func (s *DispatchRPCServer) ListDispatchOrders(ctx context.Context, req *dispatch.ListDispatchOrdersRequest) (*dispatch.ListDispatchOrdersResponse, error) {
	// 暂时返回空的订单列表
	return &dispatch.ListDispatchOrdersResponse{
		Orders: []*dispatch.DispatchOrder{},
		Total:  0,
	}, nil
}
