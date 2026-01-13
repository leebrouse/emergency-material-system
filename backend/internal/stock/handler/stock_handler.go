package handler

import (
	"net/http"
	"strconv"

	"github.com/emergency-material-system/backend/internal/stock/service"

	"github.com/gin-gonic/gin"
)

// StockHandler 物资库存处理器
type StockHandler struct {
	stockService service.StockService
}

// NewStockHandler 创建物资库存处理器
func NewStockHandler(stockService service.StockService) *StockHandler {
	return &StockHandler{
		stockService: stockService,
	}
}

// ListMaterials 获取物资列表
// GET /api/v1/stock/materials?page=1&page_size=10
func (h *StockHandler) ListMaterials(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	materials, total, err := h.stockService.ListMaterials(c.Request.Context(), page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list materials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"materials": materials,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

// CreateMaterial 创建物资
// POST /api/v1/stock/materials
func (h *StockHandler) CreateMaterial(c *gin.Context) {
	// 暂时返回未实现
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}

// GetMaterial 获取物资详情
// GET /api/v1/stock/materials/:id
func (h *StockHandler) GetMaterial(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid material ID"})
		return
	}

	material, err := h.stockService.GetMaterial(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
		return
	}

	c.JSON(http.StatusOK, material)
}

// GetInventory 获取库存信息
// GET /api/v1/stock/inventory
func (h *StockHandler) GetInventory(c *gin.Context) {
	inventory, err := h.stockService.GetInventory(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get inventory"})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

// UpdateInventory 更新库存
// PUT /api/v1/stock/inventory
func (h *StockHandler) UpdateInventory(c *gin.Context) {
	// 暂时返回未实现
	c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented"})
}
