package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenValidator() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "Authorization header required",
			})
			ctx.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "Invalid token format",
			})
			ctx.Abort()
			return
		}

		token := tokenParts[1]

		// TODO: Implement JWT validation logic here
		// For now, we'll just check if token exists
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthorized",
				"message": "Invalid or expired token",
			})
			ctx.Abort()
			return
		}

		// TODO: Add user info to context after JWT validation
		// ctx.Set("userID", claims.UserID)
		// ctx.Set("userRole", claims.Role)
		// ctx.Set("businessID", claims.BusinessID)

		ctx.Next()
	}
}

func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: Get user role from context (after JWT implementation)
		// userRole, exists := ctx.Get("userRole")
		// if !exists {
		// 	ctx.JSON(http.StatusForbidden, gin.H{
		// 		"error":   "forbidden",
		// 		"message": "Role information not found",
		// 	})
		// 	ctx.Abort()
		// 	return
		// }

		// TODO: Check if user has required role
		// if userRole != requiredRole {
		// 	ctx.JSON(http.StatusForbidden, gin.H{
		// 		"error":   "forbidden",
		// 		"message": "Insufficient permissions",
		// 	})
		// 	ctx.Abort()
		// 	return
		// }

		ctx.Next()
	}
}

func RequireBusinessAccess() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: Validate business access
		// businessID, exists := ctx.Get("businessID")
		// if !exists {
		// 	ctx.JSON(http.StatusForbidden, gin.H{
		// 		"error":   "forbidden",
		// 		"message": "Business access required",
		// 	})
		// 	ctx.Abort()
		// 	return
		// }

		ctx.Next()
	}
}
