package handler

import (
	"net/http"

	"github.com/emergency-material-system/backend/internal/statistics/service"

	"github.com/gin-gonic/gin"
)

// StatisticsHandler 统计处理器 - 实现生成的 ServerInterface
type StatisticsHandler struct {
	statisticsService service.StatisticsService
}

// NewStatisticsHandler 创建统计处理器
func NewStatisticsHandler(statisticsService service.StatisticsService) *StatisticsHandler {
	return &StatisticsHandler{
		statisticsService: statisticsService,
	}
}

// GetStatisticsReports 获取统计报表 - 实现 ServerInterface
func (h *StatisticsHandler) GetStatisticsReports(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetStatisticsSummary 获取统计汇总 - 实现 ServerInterface
func (h *StatisticsHandler) GetStatisticsSummary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetStatisticsTrends 获取统计趋势 - 实现 ServerInterface
func (h *StatisticsHandler) GetStatisticsTrends(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
