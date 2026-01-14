package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// AuthService 认证服务接口
type AuthService interface {
	Login(ctx context.Context, username, password string) (token, refreshToken string, expiresIn int64, err error)
	Logout(ctx context.Context, token string) error
	RefreshToken(ctx context.Context, refreshToken string) (token, newRefreshToken string, expiresIn int64, err error)
	ValidateToken(ctx context.Context, token string) (bool, uint, []string, error)
}

// authService 认证服务实现
type authService struct {
	jwtSecret   []byte
	tokenExpiry time.Duration
	// db          *gorm.DB
}

// NewAuthService 创建认证服务
func NewAuthService() AuthService {
	return &authService{
		jwtSecret:   []byte("your-secret-key"), // 应该从配置中读取
		tokenExpiry: 24 * time.Hour,            // 应该从配置中读取
		// db:          db,	
	}
}

// Login 用户登录
func (s *authService) Login(ctx context.Context, username, password string) (token, refreshToken string, expiresIn int64, err error) {
	// 简单的模拟验证
	if username == "" || password == "" {
		return "", "", 0, errors.New("invalid credentials")
	}

	// 模拟用户ID
	userID := uint(1)

	// 生成访问令牌
	token, err = s.generateToken(userID, username, s.tokenExpiry)
	if err != nil {
		return "", "", 0, err
	}

	// 生成刷新令牌
	refreshToken, err = s.generateToken(userID, username, 7*24*time.Hour)
	if err != nil {
		return "", "", 0, err
	}

	expiresIn = int64(s.tokenExpiry.Seconds())
	return token, refreshToken, expiresIn, nil
}

// Logout 用户登出
func (s *authService) Logout(ctx context.Context, token string) error {
	// 模拟登出，暂时只记录日志
	return nil
}

// RefreshToken 刷新访问令牌
func (s *authService) RefreshToken(ctx context.Context, refreshToken string) (token, newRefreshToken string, expiresIn int64, err error) {
	// 验证刷新令牌
	claims, err := s.validateToken(refreshToken)
	if err != nil {
		return "", "", 0, errors.New("invalid refresh token")
	}

	// 获取用户ID
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return "", "", 0, errors.New("invalid token claims")
	}

	// 模拟用户名
	username := "mock_user"

	// 生成新的访问令牌
	token, err = s.generateToken(uint(userID), username, s.tokenExpiry)
	if err != nil {
		return "", "", 0, err
	}

	// 生成新的刷新令牌
	newRefreshToken, err = s.generateToken(uint(userID), username, 7*24*time.Hour)
	if err != nil {
		return "", "", 0, err
	}

	expiresIn = int64(s.tokenExpiry.Seconds())
	return token, newRefreshToken, expiresIn, nil
}

// ValidateToken 验证令牌
func (s *authService) ValidateToken(ctx context.Context, token string) (bool, uint, []string, error) {
	claims, err := s.validateToken(token)
	if err != nil {
		return false, 0, nil, err
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return false, 0, nil, errors.New("invalid token claims")
	}

	roles := []string{"user"} // 模拟角色
	return true, uint(userID), roles, nil
}

// generateToken 生成JWT令牌
func (s *authService) generateToken(userID uint, username string, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"roles":    []string{"user"}, // 暂时硬编码
		"exp":      time.Now().Add(expiry).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

// validateToken 验证JWT令牌
func (s *authService) validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return s.jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
