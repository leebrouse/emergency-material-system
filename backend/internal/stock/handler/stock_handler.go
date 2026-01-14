package handler

import (
	"net/http"

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

// GetStockMaterials 获取物资列表 - 实现 ServerInterface
func (h *StockHandler) GetStockMaterials(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// PostStockMaterials 创建物资 - 实现 ServerInterface
func (h *StockHandler) PostStockMaterials(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetStockMaterialsId 获取物资详情 - 实现 ServerInterface
func (h *StockHandler) GetStockMaterialsId(c *gin.Context, id int) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// GetStockInventory 获取库存信息 - 实现 ServerInterface
func (h *StockHandler) GetStockInventory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
