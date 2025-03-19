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
			zap.Any("requestBody", requestBody),
		)

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		c.Next()

		latency := time.Since(startTime)
		statusCode := c.Writer.Status()
		responseBody := blw.body.String()

		switch {
		case statusCode >= 200 && statusCode < 300: // 2xx
			global.Logger.Info("Request success",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("statusCode", statusCode),
				zap.Duration("latency", latency),
				zap.Any("responseBody", responseBody),
			)

		case statusCode == 400 || statusCode == 401 || statusCode == 403: // 400, 401, 403
			global.Logger.Info("Client error",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("statusCode", statusCode),
				zap.Duration("latency", latency),
				zap.Any("responseBody", responseBody),
			)

		case statusCode == 404 || statusCode == 429: // 404, 429
			global.Logger.Warn("Client request issue",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("statusCode", statusCode),
				zap.Duration("latency", latency),
				zap.Any("responseBody", responseBody),
			)

		case statusCode >= 500: // 5xx
			global.Logger.Error("Server error",
				zap.String("method", method),
				zap.String("path", path),
				zap.Int("statusCode", statusCode),
				zap.Duration("latency", latency),
				zap.Any("responseBody", responseBody),
			)
		}
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
