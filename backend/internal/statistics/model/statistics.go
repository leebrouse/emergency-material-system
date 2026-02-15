package model

import "time"

// Summary 统计总览
type Summary struct {
	TotalMaterials    int64 `json:"total_materials"`
	TotalRequests     int64 `json:"total_requests"`
	PendingRequests   int64 `json:"pending_requests"`
	CompletedRequests int64 `json:"completed_requests"`
}

// MaterialStat 物资分类统计
type MaterialStat struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}

// RequestStat 需求状态统计
type RequestStat struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

// TrendPoint 趋势数据点
type TrendPoint struct {
	Date  string `json:"date"`
	Value int64  `json:"value"`
}

// Report 基础报表
type Report struct {
	Title     string       `json:"title"`
	Data      []TrendPoint `json:"data"`
	UpdatedAt time.Time    `json:"updated_at"`
}
