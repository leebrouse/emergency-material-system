package handler

import (
	"net/http"

	"github.com/emergency-material-system/backend/internal/auth/service"
	"github.com/emergency-material-system/backend/internal/common/genopenapi/auth"

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

// PostAuthLogin 用户登录 - 实现 ServerInterface
func (h *AuthHandler) PostAuthLogin(c *gin.Context) {
	var body auth.PostAuthLoginJSONBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if body.Username == nil || body.Password == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username and password are required"})
		return
	}

	token, refreshToken, expiresIn, err := h.authService.Login(c.Request.Context(), *body.Username, *body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  token,
		"refresh_token": refreshToken,
		"expires_in":    expiresIn,
		"token_type":    "Bearer",
	})
}

// PostAuthLogout 用户登出 - 实现 ServerInterface
func (h *AuthHandler) PostAuthLogout(c *gin.Context) {
	// 实际项目中可能需要从 Header 获取 token 并加入黑名单
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "logged out successfully",
	})
}

// PostAuthRefresh 刷新令牌 - 实现 ServerInterface
func (h *AuthHandler) PostAuthRefresh(c *gin.Context) {
	var body auth.PostAuthRefreshJSONBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	token, newRefreshToken, expiresIn, err := h.authService.RefreshToken(c.Request.Context(), body.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  token,
		"refresh_token": newRefreshToken,
		"expires_in":    expiresIn,
		"token_type":    "Bearer",
	})
}
