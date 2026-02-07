package middleware

import (
	"slices"
	"fmt"
	"net/http"
	"strings"

	"github.com/emergency-material-system/backend/internal/common/genproto/auth"
	"github.com/gin-gonic/gin"
)

// AuthMiddleware 身份认证中间件
func AuthMiddleware(authClient auth.AuthServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must be Bearer token"})
			c.Abort()
			return
		}

		token := parts[1]
		resp, err := authClient.ValidateToken(c.Request.Context(), &auth.ValidateTokenRequest{
			Token: token,
		})

		if err != nil || !resp.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		c.Set("user_id", uint(resp.UserId))
		c.Set("roles", resp.Roles)
		c.Next()
	}
}

// RoleMiddleware 角色权限中间件
func RoleMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRoles, exists := c.Get("roles")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "No roles found in context"})
			c.Abort()
			return
		}

		roles := userRoles.([]string)
		hasRole := false
		for _, required := range requiredRoles {
			if slices.Contains(roles, required) {
					hasRole = true
				}
			if hasRole {
				break
			}
		}

		if !hasRole {
			c.JSON(http.StatusForbidden, gin.H{"error": fmt.Sprintf("Missing required roles: %v", requiredRoles)})
			c.Abort()
			return
		}

		c.Next()
	}
}
