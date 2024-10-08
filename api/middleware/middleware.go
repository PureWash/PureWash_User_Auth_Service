package middleware

import (
	"fmt"
	"net/http"
	"user-service/internal/config"
	"user-service/internal/security"

	"github.com/gin-gonic/gin"
)

func CorsMiddileware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.Writer.Header().Set("Content-Type", "application/json")
        ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
        ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
        ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
        ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

        if ctx.Request.Method == "OPTIONS" {
            fmt.Println("OPTIONS request detected:", ctx.Request.URL.Path) // Qo'shimcha log
            ctx.AbortWithStatus(200)
            return
        }
        ctx.Next()
    }
}


func IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Errorf("authorization header is required"))
			return
		}

		claims, err := security.ExtractClaims(auth, config.Load().SECRET_KEY)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Errorf("invalid token claims: %s", err))
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
