package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/emergency-material-system/backend/internal/common/config"
	"github.com/emergency-material-system/backend/internal/logistics/handler"
	"github.com/emergency-material-system/backend/internal/logistics/service"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting logistics service...")

	// 暂时使用mock服务，不连接数据库
	logisticsService := service.NewMockLogisticsService()

	// 初始化处理器
	logisticsHandler := handler.NewLogisticsHandler(logisticsService)

	// 启动REST API服务器 (端口8085)
	startRESTServer(logisticsHandler)

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down logistics service...")
}

// startRESTServer 启动REST API服务器
func startRESTServer(handler *handler.LogisticsHandler) {
	// 创建Gin引擎
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "logistics"})
	})

	// API路由组
	api := r.Group("/api/v1")
	{
		logistics := api.Group("/logistics")
		{
			logistics.GET("/tracking/:id", handler.GetTracking)
			logistics.POST("/tracking", handler.CreateTracking)
			logistics.PUT("/tracking/:id", handler.UpdateTracking)
		}
	}

	fmt.Println("REST API server starting on port 8085")
	if err := r.Run(":8085"); err != nil {
		fmt.Printf("Failed to start REST server: %v\n", err)
		os.Exit(1)
	}
}
