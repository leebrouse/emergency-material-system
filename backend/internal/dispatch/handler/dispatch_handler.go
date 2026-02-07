package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/emergency-material-system/backend/internal/common/genopenapi/dispatch"
	"github.com/emergency-material-system/backend/internal/dispatch/model"
	"github.com/emergency-material-system/backend/internal/dispatch/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DispatchHandler struct {
	svc service.DispatchService
}

func NewDispatchHandler(svc service.DispatchService) *DispatchHandler {
	return &DispatchHandler{svc: svc}
}

// PostDispatchRequests 创建需求申报
func (h *DispatchHandler) PostDispatchRequests(c *gin.Context) {
	var body dispatch.CreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 防止nil pointer
	desc := ""
	if body.Description != nil {
		desc = *body.Description
	}

	req := &model.DemandRequest{
		MaterialID:  uint(body.MaterialId),
		Quantity:    int64(body.Quantity),
		Urgency:     model.UrgencyLevel(body.UrgencyLevel),
		TargetArea:  body.TargetArea,
		Description: desc,
	}

	if err := h.svc.CreateDemandRequest(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create demand request: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, req)
}

// GetDispatchRequests 获取需求单列表
func (h *DispatchHandler) GetDispatchRequests(c *gin.Context, params dispatch.GetDispatchRequestsParams) {
	page := 1
	if params.Page != nil {
		page = *params.Page
	}
	pageSize := 10
	if params.PageSize != nil {
		pageSize = *params.PageSize
	}
	status := ""
	if params.Status != nil {
		status = *params.Status
	}

	list, total, err := h.svc.ListDemandRequests(c.Request.Context(), page, pageSize, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  list,
		"total": total,
	})
}

// GetDispatchRequestsId 详情
func (h *DispatchHandler) GetDispatchRequestsId(c *gin.Context, id int) {
	req, err := h.svc.GetDemandRequest(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "request not found"})
		return
	}
	c.JSON(http.StatusOK, req)
}

// PostDispatchRequestsIdAudit 审核
func (h *DispatchHandler) PostDispatchRequestsIdAudit(c *gin.Context, id int) {
	var body dispatch.AuditRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	remark := ""
	if body.Remark != nil {
		remark = *body.Remark
	}

	if err := h.svc.AuditDemandRequest(c.Request.Context(), uint(id), string(body.Action), remark); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "demand request not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to audit request: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "audited"})
}

// GetDispatchRequestsIdAllocationSuggestion 推荐分配
func (h *DispatchHandler) GetDispatchRequestsIdAllocationSuggestion(c *gin.Context, id int) {
	suggestions, err := h.svc.SuggestAllocation(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, suggestions)
}

// PostDispatchTasks 创建任务
func (h *DispatchHandler) PostDispatchTasks(c *gin.Context) {
	var body dispatch.CreateDispatchTask
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.RequestId <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request_id"})
		return
	}
	if len(body.Allocations) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "allocations cannot be empty"})
		return
	}

	allocations := make([]service.AllocationSuggestion, len(body.Allocations))
	for i, a := range body.Allocations {
		allocations[i] = service.AllocationSuggestion{
			InventoryID: uint(a.InventoryId),
			Quantity:    int64(a.Quantity),
		}
	}

	taskID, err := h.svc.CreateDispatchTask(c.Request.Context(), uint(body.RequestId), allocations)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "demand request not found"})
			return
		}
		if strings.Contains(err.Error(), "must be approved") || strings.Contains(err.Error(), "stock lock failed") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create dispatch task: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"task_id": taskID})
}

// GetDispatchTasks 列表
func (h *DispatchHandler) GetDispatchTasks(c *gin.Context) {
	list, total, err := h.svc.ListDispatchTasks(c.Request.Context(), 1, 100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": list, "total": total})
}
