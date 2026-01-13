package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/emergency-material-system/backend/internal/common/config"
	"github.com/emergency-material-system/backend/internal/stock/handler"
	"github.com/emergency-material-system/backend/internal/stock/rpc"
	"github.com/emergency-material-system/backend/internal/stock/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Starting stock service...")

	// 暂时使用mock服务，不连接数据库
	stockService := service.NewMockStockService()

	// 初始化处理器
	stockHandler := handler.NewStockHandler(stockService)

	// 初始化gRPC服务
	stockRPCServer := rpc.NewStockRPCServer(stockService)

	// 启动gRPC服务器 (端口9092)
	go startGRPCServer(stockRPCServer)

	// 启动REST API服务器 (端口8082)
	go startRESTServer(stockHandler)

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down stock service...")
}

// startGRPCServer 启动gRPC服务器
func startGRPCServer(server *rpc.StockRPCServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", viper.GetString("services.stock.grpc")))
	if err != nil {
		fmt.Printf("Failed to listen for gRPC: %v\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()

	// 注册服务
	server.Register(grpcServer)
	reflection.Register(grpcServer)

	fmt.Println("gRPC server starting on port 9092")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve gRPC: %v\n", err)
		os.Exit(1)
	}
}

// startRESTServer 启动REST API服务器
func startRESTServer(handler *handler.StockHandler) {
	// 创建Gin引擎
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "stock"})
	})

	// API路由组
	api := r.Group("/api/v1")
	{
		stock := api.Group("/stock")
		{
			stock.GET("/materials", handler.ListMaterials)
			stock.POST("/materials", handler.CreateMaterial)
			stock.GET("/materials/:id", handler.GetMaterial)
			stock.GET("/inventory", handler.GetInventory)
			stock.PUT("/inventory", handler.UpdateInventory)
		}
	}

	fmt.Println("REST API server starting on port 8082")
	if err := r.Run(":" + viper.GetString("services.stock.rest")); err != nil {
		fmt.Printf("Failed to start REST server: %v\n", err)
		os.Exit(1)
	}
}
