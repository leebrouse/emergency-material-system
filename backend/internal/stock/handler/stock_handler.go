package handler

import (
	"log"
	"net/http"

	"github.com/emergency-material-system/backend/internal/common/genopenapi/stock"
	"github.com/emergency-material-system/backend/internal/stock/model"
	"github.com/emergency-material-system/backend/internal/stock/service"

	"github.com/gin-gonic/gin"
)

// StockHandler 物资库存处理器 - 实现生成的 ServerInterface
type StockHandler struct {
	stockService service.StockService
}

// NewStockHandler 创建物资库存处理器
func NewStockHandler(stockService service.StockService) *StockHandler {
	return &StockHandler{
		stockService: stockService,
	}
}

// GetStockMaterials 获取物资列表 (支持分页和搜索)
func (h *StockHandler) GetStockMaterials(c *gin.Context, params stock.GetStockMaterialsParams) {
	page := 1
	if params.Page != nil {
		page = *params.Page
	}
	pageSize := 10
	if params.PageSize != nil {
		pageSize = *params.PageSize
	}
	search := ""
	if params.Search != nil {
		search = *params.Search
	}

	list, total, err := h.stockService.ListMaterials(c.Request.Context(), page, pageSize, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  list,
		"total": total,
		"page":  page,
	})
}

// PostStockMaterials 创建物资
func (h *StockHandler) PostStockMaterials(c *gin.Context) {
	log.Println(">>> PostStockMaterials CALLED <<<")
	var m model.Material
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.stockService.CreateMaterial(c.Request.Context(), &m); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println(">>> Material created successfully <<<")
	c.JSON(http.StatusCreated, m)
}

// GetStockMaterialsId 获取物资详情
func (h *StockHandler) GetStockMaterialsId(c *gin.Context, id int) {
	m, err := h.stockService.GetMaterial(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "material not found"})
		return
	}
	c.JSON(http.StatusOK, m)
}

// GetStockInventory 获取库存列表
func (h *StockHandler) GetStockInventory(c *gin.Context) {
	// 简化处理分页，或从 query 获取
	list, total, err := h.stockService.ListInventory(c.Request.Context(), 1, 100)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  list,
		"total": total,
	})
}

// PostStockInbound 入库操作
func (h *StockHandler) PostStockInbound(c *gin.Context) {
	var req service.InboundRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.stockService.Inbound(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "inbound successful"})
}

// PostStockOutbound 出库操作
func (h *StockHandler) PostStockOutbound(c *gin.Context) {
	var req service.OutboundRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.stockService.Outbound(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "outbound successful"})
}

// PostStockTransfer 调拨操作
func (h *StockHandler) PostStockTransfer(c *gin.Context) {
	var req service.TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.stockService.Transfer(c.Request.Context(), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "transfer successful"})
}

// GetStockStats 获取聚合统计数据
func (h *StockHandler) GetStockStats(c *gin.Context) {
	stats, err := h.stockService.GetInventoryStats(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}
