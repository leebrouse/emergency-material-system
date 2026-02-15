package handler

import (
	"fmt"
	"net/http"

	"github.com/emergency-material-system/backend/internal/common/genopenapi/logistics"
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
	var body logistics.PostLogisticsTrackingJSONBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	description := ""
	if body.Description != nil {
		description = *body.Description
	}

	status := "created"
	if body.Status != nil {
		status = string(*body.Status)
	}

	tracking, err := h.logisticsService.CreateTracking(c.Request.Context(), uint(body.RequestId), description, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tracking)
}

// GetLogisticsTrackingId 获取物流追踪信息 - 实现 ServerInterface
func (h *LogisticsHandler) GetLogisticsTrackingId(c *gin.Context, id int) {
	tracking, err := h.logisticsService.GetTrajectory(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "tracking not found"})
		return
	}
	c.JSON(http.StatusOK, tracking)
}

// PutLogisticsTrackingId 更新物流追踪状态 - 实现 ServerInterface
func (h *LogisticsHandler) PutLogisticsTrackingId(c *gin.Context, id int) {
	var body logistics.PutLogisticsTrackingIdJSONBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	description := ""
	if body.Description != nil {
		description = *body.Description
	}

	err := h.logisticsService.UpdateTracking(c.Request.Context(), uint(id), string(body.Status), description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "updated"})
}

// RecordNodeRequest 轨迹节点请求结构
type RecordNodeRequest struct {
	Location    string  `json:"location"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Status      string  `json:"status"`
	Description string  `json:"description"`
}

// PostTrajectoryNode 记录轨迹节点 (自定义扩展接口)
func (h *LogisticsHandler) PostTrajectoryNode(c *gin.Context) {
	idStr := c.Param("id")
	var id uint
	fmt.Sscanf(idStr, "%d", &id)

	var req RecordNodeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.logisticsService.RecordTrajectoryNode(c.Request.Context(), id, req.Location, req.Latitude, req.Longitude, req.Status, req.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "node_recorded"})
}
