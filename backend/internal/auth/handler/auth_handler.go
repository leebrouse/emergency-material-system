package handler

import (
	"net/http"

	"github.com/emergency-material-system/backend/internal/auth/service"

	"github.com/gin-gonic/gin"
)

// AuthHandler 认证处理器 - 实现生成的 ServerInterface
type AuthHandler struct {
	authService service.AuthService
}

// NewAuthHandler 创建认证处理器
func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// PostAuthLogin 用户登录 - 实现 ServerInterface（简化版，只返回 200 OK）
func (h *AuthHandler) PostAuthLogin(c *gin.Context) {
	
}

// PostAuthLogout 用户登出 - 实现 ServerInterface（简化版，只返回 200 OK）
func (h *AuthHandler) PostAuthLogout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// PostAuthRefresh 刷新令牌 - 实现 ServerInterface（简化版，只返回 200 OK）
func (h *AuthHandler) PostAuthRefresh(c *gin.Context) {
	// var _ auth.PostAuthRefreshJSONRequestBody
	// _ = c.ShouldBindJSON(&_)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
