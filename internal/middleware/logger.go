package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now()
		c.Next()
		logger.Info("Request received",

			zap.String("path", c.Request.URL.Path),
			zap.String("url", c.Request.URL.String()),
			zap.String("method", c.Request.Method),
			zap.String("ip", c.ClientIP()),
			zap.String("remote_ip", c.RemoteIP()),
			zap.String("proto", c.Request.Proto),
			zap.String("host", c.Request.Host),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("referer", c.Request.Referer()),
			zap.String("content-type", c.Request.Header.Get("Content-Type")),
			zap.String("content-length", c.Request.Header.Get("Content-Length")),
			zap.String("protocol", c.Request.Proto),
			zap.Time("@timestamp", start),
			zap.Duration("duration", time.Since(start)),
			zap.Int("status", c.Writer.Status()),
		)
	}
}
