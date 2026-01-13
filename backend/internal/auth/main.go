package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/emergency-material-system/backend/internal/auth/handler"
	"github.com/emergency-material-system/backend/internal/auth/rpc"
	"github.com/emergency-material-system/backend/internal/auth/service"
	_ "github.com/emergency-material-system/backend/internal/common/config"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Starting auth service...")

	// 暂时使用mock服务，不连接数据库
	authService := service.NewMockAuthService()

	// 初始化处理器
	authHandler := handler.NewAuthHandler(authService)

	// 初始化gRPC服务
	authRPCServer := rpc.NewAuthRPCServer(authService)

	// 启动gRPC服务器 (端口9091)
	go startGRPCServer(authRPCServer)

	// 启动REST API服务器 (端口8081)
	go startRESTServer(authHandler)

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down auth service...")
}

// startGRPCServer 启动gRPC服务器
func startGRPCServer(server *rpc.AuthRPCServer) {
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Printf("Failed to listen for gRPC: %v\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()

	// 注册服务
	server.Register(grpcServer)
	reflection.Register(grpcServer)

	fmt.Println("gRPC server starting on port 9091")
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve gRPC: %v\n", err)
		os.Exit(1)
	}
}

// startRESTServer 启动REST API服务器
func startRESTServer(handler *handler.AuthHandler) {
	// 创建Gin引擎
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "auth"})
	})

	// API路由组
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", handler.Login)
			auth.POST("/logout", handler.Logout)
			auth.POST("/refresh", handler.RefreshToken)
		}
	}

	fmt.Println("REST API server starting on port 8081")
	if err := r.Run(viper.GetString("")); err != nil {
		fmt.Printf("Failed to start REST server: %v\n", err)
		os.Exit(1)
	}
}
