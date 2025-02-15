package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/bhtoan2204/gateway/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ApiLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()

		var requestBody []byte
		if c.Request.Body != nil {
			requestBody, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		global.Logger.Info("Request in",
			zap.String("method", method),
			zap.String("path", path),
			zap.String("clientIP", clientIP),
			zap.String("userAgent", userAgent),
			zap.ByteString("requestBody", requestBody),
		)

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		latency := time.Since(startTime)
		statusCode := c.Writer.Status()
		responseBody := blw.body.String()

		if c.Errors != nil && len(c.Errors) > 0 {
			global.Logger.Error("Request failed",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("statusCode", statusCode),
				zap.Duration("latency", latency),
				zap.String("errors", c.Errors.String()),
				zap.String("responseBody", responseBody),
			)
		} else {
			global.Logger.Info("Request out",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("statusCode", statusCode),
				zap.Duration("latency", latency),
				zap.String("responseBody", responseBody),
			)
		}
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
