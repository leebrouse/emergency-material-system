package rpc

import (
	"context"

	"github.com/emergency-material-system/backend/internal/auth/service"
	"github.com/emergency-material-system/backend/internal/common/genproto/auth"

	"google.golang.org/grpc"
)

// AuthRPCServer 认证gRPC服务器
type AuthRPCServer struct {
	auth.UnimplementedAuthServiceServer
	authService service.AuthService
}

// NewAuthRPCServer 创建认证gRPC服务器
func NewAuthRPCServer(authService service.AuthService) *AuthRPCServer {
	return &AuthRPCServer{
		authService: authService,
	}
}

// Register 注册gRPC服务
func (s *AuthRPCServer) Register(server *grpc.Server) {
	auth.RegisterAuthServiceServer(server, s)
}

// ValidateToken 验证令牌
func (s *AuthRPCServer) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {
	valid, userID, roles, err := s.authService.ValidateToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	return &auth.ValidateTokenResponse{
		Valid:  valid,
		UserId: int64(userID),
		Roles:  roles,
	}, nil
}
