package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/emergency-material-system/backend/internal/common/config"
	"github.com/emergency-material-system/backend/internal/common/genopenapi/dispatch"
	"github.com/emergency-material-system/backend/internal/dispatch/handler"
	"github.com/emergency-material-system/backend/internal/dispatch/rpc"
	"github.com/emergency-material-system/backend/internal/dispatch/service"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Starting dispatch service...")

	// 暂时使用mock服务，不连接数据库
	dispatchService := service.NewDispatchService()

	// 初始化处理器
	dispatchHandler := handler.NewDispatchHandler(dispatchService)

	// 初始化gRPC服务
	dispatchRPCServer := rpc.NewDispatchRPCServer(dispatchService)

	// 启动gRPC服务器 (端口9093)
	go startGRPCServer(dispatchRPCServer)

	// 启动REST API服务器 (端口8083)
	startRESTServer(dispatchHandler)

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down dispatch service...")
}

// startGRPCServer 启动gRPC服务器
func startGRPCServer(server *rpc.DispatchRPCServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", viper.GetString("services.dispatch.grpc")))
	if err != nil {
		fmt.Printf("Failed to listen for gRPC: %v\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()

	// 注册服务
	server.Register(grpcServer)
	reflection.Register(grpcServer)

	fmt.Println("gRPC server starting on port ", viper.GetString("services.dispatch.grpc"))
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve gRPC: %v\n", err)
		os.Exit(1)
	}
}

// startRESTServer 启动REST API服务器
func startRESTServer(handler *handler.DispatchHandler) {
	// 创建Gin引擎
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "dispatch"})
	})

	// API路由组
	// 使用生成的路由注册器
	api := r.Group("/api/v1")
	{
		// 使用生成的 RegisterHandlers 自动注册所有路由
		dispatch.RegisterHandlers(api, handler)
	}

	fmt.Println("REST API server starting on port ", viper.GetString("services.dispatch.rest"))
	if err := r.Run(":" + viper.GetString("services.dispatch.rest")); err != nil {
		fmt.Printf("Failed to start REST server: %v\n", err)
		os.Exit(1)
	}
}
