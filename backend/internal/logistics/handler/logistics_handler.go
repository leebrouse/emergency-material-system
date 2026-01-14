package handler

import (
	"net/http"

	"github.com/emergency-material-system/backend/internal/logistics/service"

	"github.com/gin-gonic/gin"
)

// LogisticsHandler 物流处理器 - 实现生成的 ServerInterface
type LogisticsHandler struct {
	logisticsService service.LogisticsService
}

// NewLogisticsHandler 创建物流处理器
func NewLogisticsHandler(logisticsService service.LogisticsService) *LogisticsHandler {
	return &LogisticsHandler{
		logisticsService: logisticsService,
	}
}

// PostLogisticsTracking 创建物流追踪记录 - 实现 ServerInterface
func (h *LogisticsHandler) PostLogisticsTracking(c *gin.Context) {
	// var _ logistics.PostLogisticsTrackingJSONBody
	// _ = c.ShouldBindJSON(&_)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetLogisticsTrackingId 获取物流追踪信息 - 实现 ServerInterface
func (h *LogisticsHandler) GetLogisticsTrackingId(c *gin.Context, id int) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// PutLogisticsTrackingId 更新物流追踪状态 - 实现 ServerInterface
func (h *LogisticsHandler) PutLogisticsTrackingId(c *gin.Context, id int) {
	// var _ logistics.PutLogisticsTrackingIdJSONBody
	// _ = c.ShouldBindJSON(&_)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
