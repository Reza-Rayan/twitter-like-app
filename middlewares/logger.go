package middlewares

import (
	"github.com/Reza-Rayan/twitter-like-app/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path

		logger.Log.Info("incoming request",
			zap.String("route", path),
			zap.String("method", method),
			zap.Int("status", status),
			zap.String("ip", clientIP),
			zap.Duration("latency", latency),
			zap.Time("time", start),
		)
	}
}
