package rpc

import (
	"context"
	"fmt"

	"github.com/emergency-material-system/backend/internal/common/genproto/stock"
	"github.com/emergency-material-system/backend/internal/stock/service"

	"google.golang.org/grpc"
)

// StockRPCServer 物资库存gRPC服务器
type StockRPCServer struct {
	stock.UnimplementedStockServiceServer
	stockService service.StockService
}

// NewStockRPCServer 创建物资库存gRPC服务器
func NewStockRPCServer(stockService service.StockService) *StockRPCServer {
	return &StockRPCServer{
		stockService: stockService,
	}
}

// Register 注册gRPC服务
func (s *StockRPCServer) Register(server *grpc.Server) {
	stock.RegisterStockServiceServer(server, s)
}

// ListMaterials 获取物资列表
func (s *StockRPCServer) ListMaterials(ctx context.Context, req *stock.ListMaterialsRequest) (*stock.ListMaterialsResponse, error) {
	materials, total, err := s.stockService.ListMaterials(ctx, int(req.Page), int(req.PageSize))
	if err != nil {
		return nil, err
	}

	// 转换数据格式
	var materialProtos []*stock.Material
	for _, m := range materials {
		materialProtos = append(materialProtos, &stock.Material{
			Id:          int64(m.ID),
			Name:        m.Name,
			Description: *m.Description,
			Category:    *m.Category,
			Unit:        *m.Unit,
		})
	}

	return &stock.ListMaterialsResponse{
		Materials: materialProtos,
		Total:     int32(total),
	}, nil
}

// GetMaterial 获取物资详情
func (s *StockRPCServer) GetMaterial(ctx context.Context, req *stock.GetMaterialRequest) (*stock.GetMaterialResponse, error) {
	material, err := s.stockService.GetMaterial(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &stock.GetMaterialResponse{
		Material: &stock.Material{
			Id:          int64(material.ID),
			Name:        material.Name,
			Description: *material.Description,
			Category:    *material.Category,
			Unit:        *material.Unit,
		},
	}, nil
}

// CreateMaterial 创建物资
func (s *StockRPCServer) CreateMaterial(ctx context.Context, req *stock.CreateMaterialRequest) (*stock.CreateMaterialResponse, error) {
	// 暂时返回未实现
	return nil, fmt.Errorf("not implemented")
}

// GetInventory 获取库存信息
func (s *StockRPCServer) GetInventory(ctx context.Context, req *stock.GetInventoryRequest) (*stock.GetInventoryResponse, error) {
	// 暂时返回空的库存信息
	return &stock.GetInventoryResponse{
		Inventory: &stock.Inventory{
			Id:               1,
			MaterialId:       1,
			Quantity:         1000,
			ReservedQuantity: 0,
			UpdatedAt:        0,
		},
	}, nil
}

// UpdateInventory 更新库存
func (s *StockRPCServer) UpdateInventory(ctx context.Context, req *stock.UpdateInventoryRequest) (*stock.UpdateInventoryResponse, error) {
	// 暂时返回未实现
	return nil, fmt.Errorf("not implemented")
}
