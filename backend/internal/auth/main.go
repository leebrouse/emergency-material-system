package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/emergency-material-system/backend/internal/auth/handler"
	"github.com/emergency-material-system/backend/internal/auth/model"
	"github.com/emergency-material-system/backend/internal/auth/repository"
	"github.com/emergency-material-system/backend/internal/auth/rpc"
	"github.com/emergency-material-system/backend/internal/auth/service"
	_ "github.com/emergency-material-system/backend/internal/common/config"
	"github.com/emergency-material-system/backend/internal/common/database"
	"github.com/emergency-material-system/backend/internal/common/genopenapi/auth"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Starting auth service...")

	// 1. 初始化数据库
	db := database.MustInitMySQL(
		"services.auth.mysql",
		&model.User{},
		&model.Role{},
	)

	// 2. 初始化各层
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)

	// 3. 初始化基础数据 (Roles & Admin)
	seedData(authRepo, authService)

	authHandler := handler.NewAuthHandler(authService)
	authRPCServer := rpc.NewAuthRPCServer(authService)

	// 4. 启动服务
	go startGRPCServer(authRPCServer)
	startRESTServer(authHandler)

	fmt.Println("Shutting down auth service...")
}

func seedData(repo repository.AuthRepository, svc service.AuthService) {
	ctx := context.Background()

	// 创建基础角色
	roles := []string{model.RoleAdmin, model.RoleManager, model.RoleRescue}
	for _, roleName := range roles {
		_, err := repo.GetRoleByName(ctx, roleName)
		if err != nil {
			fmt.Printf("Seeding role: %s\n", roleName)
			repo.CreateRole(ctx, &model.Role{
				Name:        roleName,
				Description: fmt.Sprintf("%s role", roleName),
			})
		}
	}

	//创建初始管理员
	adminUsername := "admin"
	_, err := repo.GetUserByUsername(ctx, adminUsername)
	if err != nil {
		fmt.Println("Seeding admin user...")
		err = svc.Register(ctx, adminUsername, "admin123", "admin@emergency.com", "13800000000", []string{model.RoleAdmin})
		if err != nil {
			fmt.Printf("Failed to seed admin user: %v\n", err)
		}
	}
}

func startGRPCServer(server *rpc.AuthRPCServer) {
	port := viper.GetString("services.auth.grpc")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("Failed to listen for gRPC: %v\n", err)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	server.Register(grpcServer)
	reflection.Register(grpcServer)

	fmt.Printf("gRPC server starting on port %s\n", port)
	if err := grpcServer.Serve(lis); err != nil {
		fmt.Printf("Failed to serve gRPC: %v\n", err)
		os.Exit(1)
	}
}

func startRESTServer(handler *handler.AuthHandler) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "service": "auth"})
	})

	api := r.Group("/api/v1")
	auth.RegisterHandlers(api, handler)

	port := viper.GetString("services.auth.rest")
	fmt.Printf("REST API server starting on port %s\n", port)
	if err := r.Run(":" + port); err != nil {
		fmt.Printf("Failed to start REST server: %v\n", err)
		os.Exit(1)
	}
}
