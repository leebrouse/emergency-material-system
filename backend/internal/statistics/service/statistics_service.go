package service

import (
	"context"
	"fmt"
	"time"

	"github.com/emergency-material-system/backend/internal/common/genproto/dispatch"
	"github.com/emergency-material-system/backend/internal/common/genproto/stock"
	"github.com/emergency-material-system/backend/internal/statistics/model"
)

// StatisticsService 统计服务接口
type StatisticsService interface {
	GetSummary(ctx context.Context) (*model.Summary, error)
	GetMaterialStats(ctx context.Context) ([]model.MaterialStat, error)
	GetRequestStats(ctx context.Context) ([]model.RequestStat, error)
	GetConsumptionTrends(ctx context.Context) ([]model.TrendPoint, error)
}

// statisticsService 统计服务实现
type statisticsService struct {
	stockClient    stock.StockServiceClient
	dispatchClient dispatch.DispatchServiceClient
}

// NewStatisticsService 创建统计服务
func NewStatisticsService(stockClient stock.StockServiceClient, dispatchClient dispatch.DispatchServiceClient) StatisticsService {
	return &statisticsService{
		stockClient:    stockClient,
		dispatchClient: dispatchClient,
	}
}

// GetSummary 获取总览统计
func (s *statisticsService) GetSummary(ctx context.Context) (*model.Summary, error) {
	// 获取物资总数
	materialRes, err := s.stockClient.ListMaterials(ctx, &stock.ListMaterialsRequest{Page: 1, PageSize: 1})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch materials: %v", err)
	}

	// 获取需求总数和状态
	demandRes, err := s.dispatchClient.ListDemands(ctx, &dispatch.ListDemandsRequest{Page: 1, PageSize: 1000})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch demands: %v", err)
	}

	var completed, pending int64
	for _, d := range demandRes.Demands {
		switch d.Status {
		case "Signed":
			completed++
		case "Pending", "Dispatching", "Shipping":
			pending++
		}
	}

	return &model.Summary{
		TotalMaterials:    int64(materialRes.Total),
		TotalRequests:     int64(demandRes.Total),
		PendingRequests:   pending,
		CompletedRequests: completed,
	}, nil
}

// GetMaterialStats 获取物资分类统计
func (s *statisticsService) GetMaterialStats(ctx context.Context) ([]model.MaterialStat, error) {
	res, err := s.stockClient.ListMaterials(ctx, &stock.ListMaterialsRequest{Page: 1, PageSize: 1000})
	if err != nil {
		return nil, err
	}

	statsMap := make(map[string]int)
	for _, m := range res.Materials {
		statsMap[m.Category]++
	}

	var result []model.MaterialStat
	for k, v := range statsMap {
		result = append(result, model.MaterialStat{
			Category: k,
			Count:    v,
		})
	}
	return result, nil
}

// GetRequestStats 获取需求状态统计
func (s *statisticsService) GetRequestStats(ctx context.Context) ([]model.RequestStat, error) {
	res, err := s.dispatchClient.ListDemands(ctx, &dispatch.ListDemandsRequest{Page: 1, PageSize: 1000})
	if err != nil {
		return nil, err
	}

	statsMap := make(map[string]int)
	for _, d := range res.Demands {
		statsMap[d.Status]++
	}

	var result []model.RequestStat
	for k, v := range statsMap {
		result = append(result, model.RequestStat{
			Status: k,
			Count:  v,
		})
	}
	return result, nil
}

// GetConsumptionTrends 获取近七天消耗趋势
func (s *statisticsService) GetConsumptionTrends(ctx context.Context) ([]model.TrendPoint, error) {
	// 获取所有出库流水
	res, err := s.stockClient.ListStockLogs(ctx, &stock.ListStockLogsRequest{Type: "Outbound", Limit: 1000})
	if err != nil {
		return nil, err
	}

	// 按日期分组汇总
	dailyMap := make(map[string]int64)
	for _, l := range res.Logs {
		dateStr := time.Unix(l.CreatedAt, 0).Format("01-02")
		dailyMap[dateStr] += -l.QuantityChange // 消耗为正值
	}

	// 生成最近7天的数据点 (即使没有数据也返回0)
	var result []model.TrendPoint
	for i := 6; i >= 0; i-- {
		date := time.Now().AddDate(0, 0, -i)
		dateStr := date.Format("01-02")

		weekday := ""
		switch date.Weekday() {
		case time.Monday:
			weekday = "周一"
		case time.Tuesday:
			weekday = "周二"
		case time.Wednesday:
			weekday = "周三"
		case time.Thursday:
			weekday = "周四"
		case time.Friday:
			weekday = "周五"
		case time.Saturday:
			weekday = "周六"
		case time.Sunday:
			weekday = "周日"
		}

		result = append(result, model.TrendPoint{
			Date:  weekday,
			Value: dailyMap[dateStr],
		})
	}

	return result, nil
}
