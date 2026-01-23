package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	_ "github.com/emergency-material-system/backend/internal/common/config"
	"github.com/emergency-material-system/backend/internal/common/genopenapi/stock"
	"github.com/emergency-material-system/backend/internal/stock/handler"
	"github.com/emergency-material-system/backend/internal/stock/model"
	"github.com/emergency-material-system/backend/internal/stock/repository"
	"github.com/emergency-material-system/backend/internal/stock/rpc"
	"github.com/emergency-material-system/backend/internal/stock/service"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting stock service...")

	// 数据库 DSN 配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("db.mysql.user"),
		viper.GetString("db.mysql.password"),
		viper.GetString("db.mysql.host"),
		viper.GetString("db.mysql.port"),
		viper.GetString("db.mysql.database"),
	)

	// 初始化数据库 (MySQL)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to MySQL database: %v\n", err)
		os.Exit(1)
	}

	// 自动迁移模型
	err = db.AutoMigrate(&model.Material{}, &model.Inventory{}, &model.StockLog{})
	if err != nil {
		fmt.Printf("Failed to auto-migrate: %v\n", err)
	}

	// 初始化各层
	stockRepo := repository.NewStockRepository(db)
	stockService := service.NewStockService(stockRepo)
	stockHandler := handler.NewStockHandler(stockService)

	// 初始化gRPC服务
	stockRPCServer := rpc.NewStockRPCServer(stockService)

	// 启动gRPC服务器
	go startGRPCServer(stockRPCServer)

	// 启动REST API服务器
	startRESTServer(stockHandler)

	fmt.Println("Shutting down stock service...")
}

// startGRPCServer 启动gRPC服务器
func startGRPCServer(server *rpc.StockRPCServer) {
	port := viper.GetString("services.stock.grpc")
	if port == "" {
		port = "9092"
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Printf("Failed to listen for gRPC: %v\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	server.Register(grpcServer)
	reflection.Register(grpcServer)

	fmt.Println("gRPC server starting on port ", port)
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve gRPC: %v\n", err)
		os.Exit(1)
	}
}

// startRESTServer 启动REST API服务器
func startRESTServer(h *handler.StockHandler) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "stock"})
	})

	// API 路由注册
	api := r.Group("/api/v1")
	{
		// 所有的 /stock/* 路由现在都已定义在 OpenAPI 中并通过 RegisterHandlers 注册
		stock.RegisterHandlers(api, h)
	}

	port := viper.GetString("services.stock.rest")
	if port == "" {
		port = "8082"
	}
	fmt.Println("REST API server starting on port ", port)
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("Failed to start REST server: %v\n", err)
		os.Exit(1)
	}
}
