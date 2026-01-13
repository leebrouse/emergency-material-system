package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	
	"github.com/emergency-material-system/backend/internal/statistics/handler"
	"github.com/emergency-material-system/backend/internal/statistics/service"
	_"github.com/emergency-material-system/backend/internal/common/config"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting statistics service...")

	// 暂时使用mock服务，不连接数据库
	statisticsService := service.NewMockStatisticsService()

	// 初始化处理器
	statisticsHandler := handler.NewStatisticsHandler(statisticsService)

	// 启动REST API服务器 (端口8084)
	startRESTServer(statisticsHandler)

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down statistics service...")
}

// startRESTServer 启动REST API服务器
func startRESTServer(handler *handler.StatisticsHandler) {
	// 创建Gin引擎
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "statistics"})
	})

	// API路由组
	api := r.Group("/api/v1")
	{
		stats := api.Group("/statistics")
		{
			stats.GET("/overview", handler.GetOverview)
			stats.GET("/materials", handler.GetMaterialStats)
			stats.GET("/requests", handler.GetRequestStats)
		}
	}

	fmt.Println("REST API server starting on port 8084")
	if err := r.Run(":8084"); err != nil {
		fmt.Printf("Failed to start REST server: %v\n", err)
		os.Exit(1)
	}
}
