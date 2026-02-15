package service

import (
	"context"
	"errors"
	"testing"

	"github.com/emergency-material-system/backend/internal/auth/model"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

// mockAuthRepository 手动实现 Mock
type mockAuthRepository struct {
	getUserFunc func(ctx context.Context, username string) (*model.User, error)
}

func (m *mockAuthRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return m.getUserFunc(ctx, username)
}
func (m *mockAuthRepository) CreateUser(ctx context.Context, user *model.User) error { return nil }
func (m *mockAuthRepository) GetRoleByName(ctx context.Context, name string) (*model.Role, error) {
	return nil, nil
}
func (m *mockAuthRepository) CreateRole(ctx context.Context, role *model.Role) error { return nil }

func TestLogin(t *testing.T) {
	// 设置测试配置
	viper.Set("services.auth.jwt.secret", "test-secret-key-for-unit-tests")
	viper.Set("services.auth.jwt.expiry_hours", 1)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	mockUser := &model.User{
		ID:       1,
		Username: "testuser",
		Password: string(hashedPassword),
		Roles: []model.Role{
			{Name: "admin"},
		},
	}

	repo := &mockAuthRepository{}
	svc := NewAuthService(repo)

	t.Run("Success", func(t *testing.T) {
		repo.getUserFunc = func(ctx context.Context, username string) (*model.User, error) {
			if username == "testuser" {
				return mockUser, nil
			}
			return nil, errors.New("not found")
		}

		token, refreshToken, expiresIn, err := svc.Login(context.Background(), "testuser", "password123")

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if token == "" {
			t.Error("expected access token to be non-empty")
		}
		if refreshToken == "" {
			t.Error("expected refresh token to be non-empty")
		}
		if expiresIn != 3600 {
			t.Errorf("expected expiresIn 3600, got %d", expiresIn)
		}
	})

	t.Run("InvalidPassword", func(t *testing.T) {
		repo.getUserFunc = func(ctx context.Context, username string) (*model.User, error) {
			return mockUser, nil
		}

		_, _, _, err := svc.Login(context.Background(), "testuser", "wrongpassword")

		if err == nil {
			t.Fatal("expected error for invalid password, got nil")
		}
		if err.Error() != "invalid credentials" {
			t.Errorf("expected 'invalid credentials' error, got %v", err)
		}
	})

	t.Run("UserNotFound", func(t *testing.T) {
		repo.getUserFunc = func(ctx context.Context, username string) (*model.User, error) {
			return nil, errors.New("record not found")
		}

		_, _, _, err := svc.Login(context.Background(), "nonexistent", "password123")

		if err == nil {
			t.Fatal("expected error for non-existent user, got nil")
		}
		if err.Error() != "invalid credentials" {
			t.Errorf("expected 'invalid credentials' error, got %v", err)
		}
	})
}
