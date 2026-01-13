package handler

import (
	"net/http"
	"strconv"

	"github.com/emergency-material-system/backend/internal/common/genopenapi/dispatch"
	"github.com/emergency-material-system/backend/internal/dispatch/service"

	"github.com/gin-gonic/gin"
)

// DispatchHandler 调度处理器
type DispatchHandler struct {
	dispatchService service.DispatchService
}

// NewDispatchHandler 创建调度处理器
func NewDispatchHandler(dispatchService service.DispatchService) *DispatchHandler {
	return &DispatchHandler{
		dispatchService: dispatchService,
	}
}

// ListRequests 获取需求申报列表
func (h *DispatchHandler) ListRequests(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	requests, total, err := h.dispatchService.ListRequests(c.Request.Context(), page, pageSize)
	if err != nil {
		// h.logger.Error("Failed to list requests", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list requests"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"requests":  requests,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// CreateRequest 创建需求申报
func (h *DispatchHandler) CreateRequest(c *gin.Context) {
	var req dispatch.PostDispatchRequestsJSONBody
	if err := c.ShouldBindJSON(&req); err != nil {
		// h.logger.Error("Failed to bind create request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	request, err := h.dispatchService.CreateRequest(c.Request.Context(), req)
	if err != nil {
		// h.logger.Error("Failed to create request", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	c.JSON(http.StatusCreated, request)
}

// GetRequest 获取需求申报详情
func (h *DispatchHandler) GetRequest(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}

	request, err := h.dispatchService.GetRequest(c.Request.Context(), uint(id))
	if err != nil {
		// h.logger.Error("Failed to get request", zap.Uint("id", uint(id)), zap.Error(err))
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}

	c.JSON(http.StatusOK, request)
}

// UpdateRequestStatus 更新需求申报状态
func (h *DispatchHandler) UpdateRequestStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}

	var req dispatch.PutDispatchRequestsIdStatusJSONBody
	if err := c.ShouldBindJSON(&req); err != nil {
		// h.logger.Error("Failed to bind update status request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err = h.dispatchService.UpdateRequestStatus(c.Request.Context(), uint(id), req)
	if err != nil {
		// h.logger.Error("Failed to update request status", zap.Uint("id", uint(id)), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update request status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Request status updated successfully"})
}
