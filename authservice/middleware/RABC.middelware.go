package middleware

import (
	"net/http"
	"strings"

	"auth/domain"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// RoleBasedAccess checks if the user has one of the allowed roles
func RoleBasedAccess(secret string, allowedRoles []domain.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header missing"})
			c.Abort()
			return
		}

		// Parse the token
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenStr, &domain.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			c.Abort()
			return
		}

		// Validate role in claims
		claims, ok := token.Claims.(*domain.UserClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		// Check if the user's role matches any of the allowed roles
		for _, role := range allowedRoles {
			if claims.Role == string(role) {
				c.Next()
				return
			}
		}

		// Role not allowed
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		c.Abort()
	}
}
