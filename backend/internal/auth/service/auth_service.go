package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/emergency-material-system/backend/internal/auth/model"
	"github.com/emergency-material-system/backend/internal/auth/repository"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// AuthService 认证服务接口
type AuthService interface {
	Login(ctx context.Context, username, password string) (token, refreshToken string, expiresIn int64, err error)
	Logout(ctx context.Context, token string) error
	RefreshToken(ctx context.Context, refreshToken string) (token, newRefreshToken string, expiresIn int64, err error)
	ValidateToken(ctx context.Context, token string) (bool, uint, []string, error)
	Register(ctx context.Context, username, password, email, phone string, roles []string) error
}

// authService 认证服务实现
type authService struct {
	repo        repository.AuthRepository
	jwtSecret   []byte
	tokenExpiry time.Duration
}

// NewAuthService 创建认证服务
func NewAuthService(repo repository.AuthRepository) AuthService {
	secret := viper.GetString("services.auth.jwt.secret")
	if secret == "" {
		secret = "default-secret-key"
	}
	expiry := viper.GetInt("services.auth.jwt.expiry_hours")
	if expiry == 0 {
		expiry = 24
	}

	return &authService{
		repo:        repo,
		jwtSecret:   []byte(secret),
		tokenExpiry: time.Duration(expiry) * time.Hour,
	}
}

// Login 用户登录
func (s *authService) Login(ctx context.Context, username, password string) (token, refreshToken string, expiresIn int64, err error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return "", "", 0, errors.New("invalid credentials")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", "", 0, errors.New("invalid credentials")
	}

	var roleNames []string
	for _, role := range user.Roles {
		roleNames = append(roleNames, role.Name)
	}

	// 生成访问令牌
	token, err = s.generateToken(user.ID, user.Username, roleNames, s.tokenExpiry)
	if err != nil {
		return "", "", 0, err
	}

	// 生成刷新令牌 (7天有效期)
	refreshToken, err = s.generateToken(user.ID, user.Username, roleNames, 7*24*time.Hour)
	if err != nil {
		return "", "", 0, err
	}

	expiresIn = int64(s.tokenExpiry.Seconds())
	return token, refreshToken, expiresIn, nil
}

// Register 用户注册 (内部使用或管理员使用)
func (s *authService) Register(ctx context.Context, username, password, email, phone string, roleNames []string) error {
	// 检查用户是否已存在
	_, err := s.repo.GetUserByUsername(ctx, username)
	if err == nil {
		return errors.New("user already exists")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
		Phone:    phone,
		Status:   model.UserStatusActive,
	}

	// 分配角色
	for _, name := range roleNames {
		role, err := s.repo.GetRoleByName(ctx, name)
		if err != nil {
			return fmt.Errorf("role %s does not exist", name)
		}
		user.Roles = append(user.Roles, *role)
	}

	return s.repo.CreateUser(ctx, user)
}

// Logout 用户登出
func (s *authService) Logout(ctx context.Context, token string) error {
	// 可以实现黑名单机制，暂时简单返回
	return nil
}

// RefreshToken 刷新访问令牌
func (s *authService) RefreshToken(ctx context.Context, refreshToken string) (token, newRefreshToken string, expiresIn int64, err error) {
	claims, err := s.validateToken(refreshToken)
	if err != nil {
		return "", "", 0, errors.New("invalid refresh token")
	}

	userID := uint(claims["user_id"].(float64))
	username := claims["username"].(string)
	rolesInterface := claims["roles"].([]interface{})
	var roles []string
	for _, r := range rolesInterface {
		roles = append(roles, r.(string))
	}

	token, err = s.generateToken(userID, username, roles, s.tokenExpiry)
	if err != nil {
		return "", "", 0, err
	}

	newRefreshToken, err = s.generateToken(userID, username, roles, 7*24*time.Hour)
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

	userID := uint(claims["user_id"].(float64))
	rolesInterface, ok := claims["roles"].([]interface{})
	var roles []string
	if ok {
		for _, r := range rolesInterface {
			roles = append(roles, r.(string))
		}
	}

	return true, userID, roles, nil
}

func (s *authService) generateToken(userID uint, username string, roles []string, expiry time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"roles":    roles,
		"exp":      time.Now().Add(expiry).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

func (s *authService) validateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
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
