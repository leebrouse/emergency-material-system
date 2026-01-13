package handler

import (
	"net/http"

	"github.com/emergency-material-system/backend/internal/statistics/service"

	"github.com/gin-gonic/gin"
)

// StatisticsHandler 统计处理器
type StatisticsHandler struct {
	statisticsService service.StatisticsService
}

// NewStatisticsHandler 创建统计处理器
func NewStatisticsHandler(statisticsService service.StatisticsService) *StatisticsHandler {
	return &StatisticsHandler{
		statisticsService: statisticsService,
	}
}

// GetOverview 获取总览统计
func (h *StatisticsHandler) GetOverview(c *gin.Context) {
	overview, err := h.statisticsService.GetOverview(c.Request.Context())
	if err != nil {
		// h.logger.Error("Failed to get overview", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get overview"})
		return
	}

	c.JSON(http.StatusOK, overview)
}

// GetMaterialStats 获取物资统计
func (h *StatisticsHandler) GetMaterialStats(c *gin.Context) {
	stats, err := h.statisticsService.GetMaterialStats(c.Request.Context())
	if err != nil {
		// h.logger.Error("Failed to get material stats", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get material stats"})
		return
	}

	c.JSON(http.StatusOK, stats)
}

// GetRequestStats 获取需求统计
func (h *StatisticsHandler) GetRequestStats(c *gin.Context) {
	stats, err := h.statisticsService.GetRequestStats(c.Request.Context())
	if err != nil {
		// h.logger.Error("Failed to get request stats", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get request stats"})
		return
	}

	c.JSON(http.StatusOK, stats)
}
