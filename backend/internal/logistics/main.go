package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/emergency-material-system/backend/internal/common/config"
	"github.com/emergency-material-system/backend/internal/common/genopenapi/logistics" // 导入生成的包
	"github.com/emergency-material-system/backend/internal/logistics/handler"
	"github.com/emergency-material-system/backend/internal/logistics/rpc"
	"github.com/emergency-material-system/backend/internal/logistics/service"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting logistics service...")

	// 暂时使用mock服务，不连接数据库
	logisticsService := service.NewLogisticsService()

	// 初始化处理器
	logisticsHandler := handler.NewLogisticsHandler(logisticsService)

	// 初始化gRPC服务
	logisticsRPCServer := rpc.NewLogisticsRPCServer(logisticsService)

	// 启动gRPC服务器 (端口9095)
	go startGRPCServer(logisticsRPCServer)

	// 启动REST API服务器 (端口8085)
	go startRESTServer(logisticsHandler)

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down logistics service...")
}

// startGRPCServer 启动gRPC服务器
func startGRPCServer(server *rpc.LogisticsRPCServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", viper.GetString("services.logistics.grpc")))
	if err != nil {
		fmt.Printf("Failed to listen for gRPC: %v\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()

	// 注册服务
	server.Register(grpcServer)
	reflection.Register(grpcServer)

	fmt.Println("gRPC server starting on port ", viper.GetString("services.logistics.grpc"))
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve gRPC: %v\n", err)
		os.Exit(1)
	}
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

	// 使用生成的路由注册器
	api := r.Group("/api/v1")
	{
		// 使用生成的 RegisterHandlers 自动注册所有路由
		logistics.RegisterHandlers(api, handler)
	}

	fmt.Println("REST API server starting on port ", viper.GetString("services.logistics.rest"))
	if err := r.Run(":" + viper.GetString("services.logistics.rest")); err != nil {
		fmt.Printf("Failed to start REST server: %v\n", err)
		os.Exit(1)
	}
}
