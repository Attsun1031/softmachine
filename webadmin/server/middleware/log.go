package middleware

import (
	"time"

	"github.com/Attsun1031/jobnetes/utils/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	clientIP := c.ClientIP()
	method := c.Request.Method

	// Process request
	c.Next()

	// Logging
	end := time.Now()
	latency := end.Sub(start)
	statusCode := c.Writer.Status()
	logger := log.Logger.WithFields(logrus.Fields{
		"path":    path,
		"query":   raw,
		"method":  method,
		"IP":      clientIP,
		"latency": latency,
		"status":  statusCode,
	})
	if statusCode < 400 {
		logger.Info("")
	} else if statusCode < 500 {
		logger.Warn("Client error")
	} else {
		logger.Error("Server error")
	}
}
