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

// Login 用户登录
func (s *AuthRPCServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	token, refreshToken, expiresIn, err := s.authService.Login(ctx, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &auth.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}, nil
}

// Logout 用户登出
func (s *AuthRPCServer) Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	err := s.authService.Logout(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	return &auth.LogoutResponse{Success: true}, nil
}

// RefreshToken 刷新令牌
func (s *AuthRPCServer) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	token, refreshToken, expiresIn, err := s.authService.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &auth.RefreshTokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}, nil
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
