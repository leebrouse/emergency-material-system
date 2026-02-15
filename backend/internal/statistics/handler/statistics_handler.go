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
	stats, err := h.statisticsService.GetMaterialStats(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}

// GetStatisticsSummary 获取统计汇总 - 实现 ServerInterface
func (h *StatisticsHandler) GetStatisticsSummary(c *gin.Context) {
	summary, err := h.statisticsService.GetSummary(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, summary)
}

// GetStatisticsTrends 获取统计趋势 - 实现 ServerInterface
func (h *StatisticsHandler) GetStatisticsTrends(c *gin.Context) {
	trends, err := h.statisticsService.GetConsumptionTrends(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trends)
}
