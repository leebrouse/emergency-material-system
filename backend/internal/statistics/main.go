package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/emergency-material-system/backend/internal/common/config"
	"github.com/emergency-material-system/backend/internal/common/genopenapi/statistics"
	"github.com/emergency-material-system/backend/internal/statistics/handler"
	"github.com/emergency-material-system/backend/internal/statistics/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("Starting statistics service...")

	// 暂时使用mock服务，不连接数据库
	statisticsService := service.NewStatisticsService()

	// 初始化处理器
	statisticsHandler := handler.NewStatisticsHandler(statisticsService)

	// 启动REST API服务器 (端口8084)
	startRESTServer(statisticsHandler)

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
		// 使用生成的 RegisterHandlers 自动注册所有路由
		statistics.RegisterHandlers(api, handler)
	}

	fmt.Println("REST API server starting on port ", viper.GetString("services.statistics.rest"))
	if err := r.Run(":" + viper.GetString("services.statistics.rest")); err != nil {
		fmt.Printf("Failed to start REST server: %v\n", err)
		os.Exit(1)
	}
}
