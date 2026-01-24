package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	_ "github.com/emergency-material-system/backend/internal/common/config"
	"github.com/emergency-material-system/backend/internal/common/genopenapi/dispatch"
	"github.com/emergency-material-system/backend/internal/common/genproto/stock"
	"github.com/emergency-material-system/backend/internal/dispatch/handler"
	"github.com/emergency-material-system/backend/internal/dispatch/model"
	"github.com/emergency-material-system/backend/internal/dispatch/repository"
	"github.com/emergency-material-system/backend/internal/dispatch/rpc"
	"github.com/emergency-material-system/backend/internal/dispatch/service"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting dispatch service...")

	// 1. DB Init
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("services.dispatch.mysql.user"),
		viper.GetString("services.dispatch.mysql.password"),
		viper.GetString("services.dispatch.mysql.host"),
		viper.GetString("services.dispatch.mysql.port"),
		viper.GetString("services.dispatch.mysql.database"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Failed to connect to MySQL: %v\n", err)
		os.Exit(1)
	}
	db.AutoMigrate(&model.DemandRequest{}, &model.DispatchTask{}, &model.DispatchLog{})

	// 2. gRPC Client Init (Stock Service)
	stockHost := viper.GetString("services.stock.host")
	if stockHost == "" {
		stockHost = "stock"
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

	// 3. Service & Handler Init
	dispatchRepo := repository.NewDispatchRepository(db)
	dispatchService := service.NewDispatchService(dispatchRepo, stockClient)
	dispatchHandler := handler.NewDispatchHandler(dispatchService)

	// 4. gRPC Server Init (Dispatch Service)
	dispatchRPCServer := rpc.NewDispatchRPCServer(dispatchService)

	go startGRPCServer(dispatchRPCServer)
	startRESTServer(dispatchHandler)
}

func startGRPCServer(server *rpc.DispatchRPCServer) {
	port := viper.GetString("services.dispatch.grpc")
	if port == "" {
		port = "9093"
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Printf("Failed to listen for gRPC: %v\n", err)
		return
	}

	grpcServer := grpc.NewServer()
	server.Register(grpcServer)
	reflection.Register(grpcServer)

	fmt.Println("gRPC server starting on port ", port)
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve gRPC: %v\n", err)
	}
}

// REST API Server
func startRESTServer(h *handler.DispatchHandler) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "dispatch"})
	})

	api := r.Group("/api/v1")
	{
		dispatch.RegisterHandlers(api, h)
	}

	port := viper.GetString("services.dispatch.rest")
	if port == "" {
		port = "8083"
	}
	fmt.Println("REST API server starting on port ", port)
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("Failed to start REST server: %v\n", err)
	}
}
