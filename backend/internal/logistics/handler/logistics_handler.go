package handler

import (
	"net/http"
	"strconv"

	"github.com/emergency-material-system/backend/internal/logistics/service"

	"github.com/gin-gonic/gin"
)

// LogisticsHandler 物流处理器
type LogisticsHandler struct {
	logisticsService service.LogisticsService
}

// NewLogisticsHandler 创建物流处理器
func NewLogisticsHandler(logisticsService service.LogisticsService) *LogisticsHandler {
	return &LogisticsHandler{
		logisticsService: logisticsService,
	}
}

// GetTracking 获取物流追踪信息
func (h *LogisticsHandler) GetTracking(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tracking ID"})
		return
	}

	tracking, err := h.logisticsService.GetTracking(c.Request.Context(), uint(id))
	if err != nil {
		// h.logger.Error("Failed to get tracking", zap.Uint("id", uint(id)), zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "Tracking not found"})
		return
	}

	c.JSON(http.StatusOK, tracking)
}

// CreateTracking 创建物流追踪记录
func (h *LogisticsHandler) CreateTracking(c *gin.Context) {
	var req struct {
		RequestID   uint   `json:"request_id" binding:"required"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		// h.logger.Error("Failed to bind create tracking request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	tracking, err := h.logisticsService.CreateTracking(c.Request.Context(), req.RequestID, req.Description, req.Status)
	if err != nil {
		// h.logger.Error("Failed to create tracking", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tracking"})
		return
	}

	c.JSON(http.StatusCreated, tracking)
}

// UpdateTracking 更新物流追踪状态
func (h *LogisticsHandler) UpdateTracking(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tracking ID"})
		return
	}

	var req struct {
		Status      string `json:"status" binding:"required"`
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		// h.logger.Error("Failed to bind update tracking request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.logisticsService.UpdateTracking(c.Request.Context(), uint(id), req.Status, req.Description)
	if err != nil {
		// h.logger.Error("Failed to update tracking", zap.Uint("id", uint(id)), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tracking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tracking updated successfully"})
}
