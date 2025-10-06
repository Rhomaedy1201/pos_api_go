package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] %s %s %d %s %s\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
		)
	})
}

func CustomLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery

		ctx.Next()

		latency := time.Since(start)
		statusCode := ctx.Writer.Status()
		method := ctx.Request.Method
		clientIP := ctx.ClientIP()

		if raw != "" {
			path = path + "?" + raw
		}

		fmt.Printf("API Request: %s %s %s - Status: %d - Duration: %v\n",
			clientIP, method, path, statusCode, latency)
	}
}

func ErrorLogger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		// Log any errors that occurred during request processing
		for _, err := range ctx.Errors {
			fmt.Printf("Error: %s\n", err.Error())
		}
	}
}
