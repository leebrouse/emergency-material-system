package service

import (
	"context"
)

// StatisticsService 统计服务接口
type StatisticsService interface {
	GetOverview(ctx context.Context) (*OverviewData, error)
	GetMaterialStats(ctx context.Context) (interface{}, error)
	GetRequestStats(ctx context.Context) (interface{}, error)
}

// statisticsService 统计服务实现
type statisticsService struct{}

// NewStatisticsService 创建统计服务
func NewStatisticsService() StatisticsService {
	return &statisticsService{}
}

// OverviewData 总览数据
type OverviewData struct {
	TotalMaterials    int64 `json:"total_materials"`
	TotalRequests     int64 `json:"total_requests"`
	PendingRequests   int64 `json:"pending_requests"`
	CompletedRequests int64 `json:"completed_requests"`
}

// GetOverview 获取总览统计
func (s *statisticsService) GetOverview(ctx context.Context) (*OverviewData, error) {

	// 返回模拟数据
	return &OverviewData{
		TotalMaterials:    10,
		TotalRequests:     25,
		PendingRequests:   5,
		CompletedRequests: 20,
	}, nil
}

// GetMaterialStats 获取物资统计
func (s *statisticsService) GetMaterialStats(ctx context.Context) (interface{}, error) {

	// 返回模拟数据
	return []map[string]interface{}{
		{"category": "医疗物资", "count": 5},
		{"category": "防护用品", "count": 3},
		{"category": "药品", "count": 2},
	}, nil
}

// GetRequestStats 获取需求统计
func (s *statisticsService) GetRequestStats(ctx context.Context) (interface{}, error) {

	// 返回模拟数据
	return []map[string]interface{}{
		{"status": "pending", "count": 5},
		{"status": "approved", "count": 10},
		{"status": "completed", "count": 8},
		{"status": "rejected", "count": 2},
	}, nil
}
