package rpc

import (
	"context"
	"fmt"
	"time"

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
	materials, total, err := s.stockService.ListMaterials(ctx, int(req.Page), int(req.PageSize), req.Keyword)
	if err != nil {
		return nil, err
	}

	var materialProtos []*stock.Material
	for _, m := range materials {
		materialProtos = append(materialProtos, &stock.Material{
			Id:          int64(m.ID),
			Name:        m.Name,
			Description: m.Description,
			Category:    m.Category,
			Unit:        m.Unit,
			CreatedAt:   m.CreatedAt.Unix(),
			UpdatedAt:   m.UpdatedAt.Unix(),
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
			Description: material.Description,
			Category:    material.Category,
			Unit:        material.Unit,
			CreatedAt:   material.CreatedAt.Unix(),
			UpdatedAt:   material.UpdatedAt.Unix(),
		},
	}, nil
}

// GetInventory 获取单个汇总库存
func (s *StockRPCServer) GetInventory(ctx context.Context, req *stock.GetInventoryRequest) (*stock.GetInventoryResponse, error) {
	// 汇总逻辑
	items, err := s.stockService.ListInventoryByMaterial(ctx, uint(req.MaterialId))
	if err != nil {
		return nil, err
	}

	var totalQty, lockedQty int64
	for _, item := range items {
		totalQty += item.Quantity
		lockedQty += item.LockedQuantity
	}

	return &stock.GetInventoryResponse{
		Inventory: &stock.Inventory{
			MaterialId:       req.MaterialId,
			Quantity:         totalQty,
			ReservedQuantity: lockedQty,
			UpdatedAt:        time.Now().Unix(),
		},
	}, nil
}

// ListInventoryItems 获取明细库存
func (s *StockRPCServer) ListInventoryItems(ctx context.Context, req *stock.ListInventoryItemsRequest) (*stock.ListInventoryItemsResponse, error) {
	items, err := s.stockService.ListInventoryByMaterial(ctx, uint(req.MaterialId))
	if err != nil {
		return nil, err
	}

	var protos []*stock.InventoryItem
	for _, item := range items {
		expiry := int64(0)
		if item.Material.ExpiryDate != nil {
			expiry = item.Material.ExpiryDate.Unix()
		}
		protos = append(protos, &stock.InventoryItem{
			Id:             int64(item.ID),
			MaterialId:     int64(item.MaterialID),
			Location:       item.WarehouseLocation,
			Quantity:       item.Quantity,
			LockedQuantity: item.LockedQuantity,
			BatchNum:       item.Material.BatchNumber,
			ExpiryDate:     expiry,
		})
	}

	return &stock.ListInventoryItemsResponse{Items: protos}, nil
}

// LockStock 锁定库存
func (s *StockRPCServer) LockStock(ctx context.Context, req *stock.LockStockRequest) (*stock.LockStockResponse, error) {
	lockItems := make(map[uint]int64)
	for _, item := range req.Items {
		lockItems[uint(item.InventoryId)] = item.Quantity
	}

	err := s.stockService.LockStock(ctx, uint(req.RequestId), lockItems)
	if err != nil {
		return &stock.LockStockResponse{Success: false, Message: err.Error()}, nil
	}

	return &stock.LockStockResponse{Success: true}, nil
}

// Unimplemented methods
func (s *StockRPCServer) CreateMaterial(ctx context.Context, req *stock.CreateMaterialRequest) (*stock.CreateMaterialResponse, error) {
	return nil, fmt.Errorf("not implemented via gRPC")
}

func (s *StockRPCServer) UpdateMaterial(ctx context.Context, req *stock.UpdateMaterialRequest) (*stock.UpdateMaterialResponse, error) {
	return nil, fmt.Errorf("not implemented via gRPC")
}

func (s *StockRPCServer) UpdateInventory(ctx context.Context, req *stock.UpdateInventoryRequest) (*stock.UpdateInventoryResponse, error) {
	return nil, fmt.Errorf("not implemented via gRPC")
}
