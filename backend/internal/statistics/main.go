package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/emergency-material-system/backend/internal/common/config"
	"github.com/emergency-material-system/backend/internal/common/genopenapi/statistics"
	"github.com/emergency-material-system/backend/internal/common/genproto/dispatch"
	"github.com/emergency-material-system/backend/internal/common/genproto/stock"
	"github.com/emergency-material-system/backend/internal/statistics/handler"
	"github.com/emergency-material-system/backend/internal/statistics/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("Starting statistics service...")

	// 1. DB Init (虽目前由于聚合可能暂无自建表，但为架构统一和后续扩展初始化)
	// database.MustInitMySQL("services.statistics.mysql")

	// 2. gRPC Client Init (Stock Service)
	stockHost := viper.GetString("services.stock.host")
	if stockHost == "" {
		stockHost = "localhost"
	}
	stockPort := viper.GetString("services.stock.grpc")
	if stockPort == "" {
		stockPort = "9092"
	}
	stockConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", stockHost, stockPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect to stock service: %v\n", err)
		os.Exit(1)
	}
	stockClient := stock.NewStockServiceClient(stockConn)

	// 3. gRPC Client Init (Dispatch Service)
	dispatchHost := viper.GetString("services.dispatch.host")
	if dispatchHost == "" {
		dispatchHost = "localhost"
	}
	dispatchPort := viper.GetString("services.dispatch.grpc")
	if dispatchPort == "" {
		dispatchPort = "9093"
	}
	dispatchConn, err := grpc.NewClient(fmt.Sprintf("%s:%s", dispatchHost, dispatchPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("Failed to connect to dispatch service: %v\n", err)
		os.Exit(1)
	}
	dispatchClient := dispatch.NewDispatchServiceClient(dispatchConn)

	// 初始化服务
	statisticsService := service.NewStatisticsService(stockClient, dispatchClient)

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
