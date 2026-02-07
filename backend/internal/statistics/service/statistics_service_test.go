package service

import (
	"context"
	"testing"

	"github.com/emergency-material-system/backend/internal/common/genproto/dispatch"
	"github.com/emergency-material-system/backend/internal/common/genproto/stock"
	"google.golang.org/grpc"
)

// Mock clients for benchmarking (simulating fast responses)
type mockStockClient struct{ stock.StockServiceClient }

func (m *mockStockClient) ListMaterials(ctx context.Context, in *stock.ListMaterialsRequest, opts ...grpc.CallOption) (*stock.ListMaterialsResponse, error) {
	materials := []*stock.Material{
		{Category: "医疗物资"},
		{Category: "防护用品"},
	}
	return &stock.ListMaterialsResponse{Total: 2, Materials: materials}, nil
}

type mockDispatchClient struct{ dispatch.DispatchServiceClient }

func (m *mockDispatchClient) ListDemands(ctx context.Context, in *dispatch.ListDemandsRequest, opts ...grpc.CallOption) (*dispatch.ListDemandsResponse, error) {
	demands := []*dispatch.Demand{
		{Status: "Signed"},
		{Status: "Pending"},
	}
	return &dispatch.ListDemandsResponse{Total: 2, Demands: demands}, nil
}

func BenchmarkGetSummary(b *testing.B) {
	s := NewStatisticsService(&mockStockClient{}, &mockDispatchClient{})
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = s.GetSummary(ctx)
	}
}

func BenchmarkGetMaterialStats(b *testing.B) {
	s := NewStatisticsService(&mockStockClient{}, &mockDispatchClient{})
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = s.GetMaterialStats(ctx)
	}
}
