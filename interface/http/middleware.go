package http

import (
	"time"

	"template/service/global_variable"
	"template/service/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		// Log Request
		latency := time.Since(t)
		requestId := c.GetString("request_id")
		logger.Logger.Infow(path,
			zap.String("request-id", requestId),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("latency", latency),
		)
	}
}

func SetRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, exist := c.Get(global_variable.KEY_REQUEST_ID); !exist {
			c.Set(global_variable.KEY_REQUEST_ID, uuid.New().String())
		}

		requestId, _ := c.Get(global_variable.KEY_REQUEST_ID)
		c.Set(global_variable.KEY_LOGGER, logger.Logger.With(
			global_variable.KEY_REQUEST_ID, requestId,
			global_variable.KEY_PART, "interface",
		))

		c.Next()
	}
}
