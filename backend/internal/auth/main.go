package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/emergency-material-system/backend/internal/auth/handler"
	"github.com/emergency-material-system/backend/internal/auth/rpc"
	"github.com/emergency-material-system/backend/internal/auth/service"
	_ "github.com/emergency-material-system/backend/internal/common/config"
	"github.com/emergency-material-system/backend/internal/common/genopenapi/auth" // 导入生成的包
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Starting auth service...")

	authService := service.NewAuthService()

	// 初始化处理器
	authHandler := handler.NewAuthHandler(authService)

	// 初始化gRPC服务
	authRPCServer := rpc.NewAuthRPCServer(authService)

	// 启动gRPC服务器 (端口9091)
	go startGRPCServer(authRPCServer)

	// 启动REST API服务器 (端口8081)
	startRESTServer(authHandler)

	fmt.Println("Shutting down auth service...")
}

// startGRPCServer 启动gRPC服务器
func startGRPCServer(server *rpc.AuthRPCServer) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", viper.GetString("services.auth.grpc")))
	if err != nil {
		fmt.Printf("Failed to listen for gRPC: %v\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()

	// 注册服务
	server.Register(grpcServer)
	reflection.Register(grpcServer)

	fmt.Println("gRPC server starting on port ", viper.GetString("services.auth.grpc"))
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

	// 使用生成的路由注册器
	api := r.Group("/api/v1")
	{
		// 使用生成的 RegisterHandlers 自动注册所有路由
		auth.RegisterHandlers(api, handler)
	}

	fmt.Println("REST API server starting on port ", viper.GetString("services.auth.rest"))
	if err := r.Run(":" + viper.GetString("services.auth.rest")); err != nil {
		fmt.Printf("Failed to start REST server: %v\n", err)
		os.Exit(1)
	}
}
