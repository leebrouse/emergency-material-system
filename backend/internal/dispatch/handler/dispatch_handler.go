package handler

import (
	"net/http"

	"github.com/emergency-material-system/backend/internal/common/genopenapi/dispatch"
	"github.com/emergency-material-system/backend/internal/dispatch/service"

	"github.com/gin-gonic/gin"
)

// DispatchHandler 调度处理器 - 实现生成的 ServerInterface
type DispatchHandler struct {
	dispatchService service.DispatchService
}

// NewDispatchHandler 创建调度处理器
func NewDispatchHandler(dispatchService service.DispatchService) *DispatchHandler {
	return &DispatchHandler{
		dispatchService: dispatchService,
	}
}

// GetDispatchRequests 获取需求申报列表 - 实现 ServerInterface
func (h *DispatchHandler) GetDispatchRequests(c *gin.Context, params dispatch.GetDispatchRequestsParams) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// PostDispatchRequests 创建需求申报 - 实现 ServerInterface
func (h *DispatchHandler) PostDispatchRequests(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetDispatchRequestsId 获取需求申报详情 - 实现 ServerInterface
func (h *DispatchHandler) GetDispatchRequestsId(c *gin.Context, id int) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// PutDispatchRequestsIdStatus 更新需求申报状态 - 实现 ServerInterface
func (h *DispatchHandler) PutDispatchRequestsId(c *gin.Context, id int) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
