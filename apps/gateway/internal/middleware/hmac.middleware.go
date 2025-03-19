package middleware

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/bhtoan2204/gateway/global"
	"github.com/bhtoan2204/gateway/pkg/response"
	"github.com/gin-gonic/gin"
)

const RedisNonceKey = "nonce:"
const RedisNonceTTL = 5 * time.Minute
const NonceTTL = 5 * time.Minute

func GenerateHMACSignature(message string) string {
	h := hmac.New(sha256.New, []byte(global.Config.SecurityConfig.HMACSecret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

func ValidateNonce(nonce string) bool {
	ctx := context.Background()
	redisKey := RedisNonceKey + nonce
	fmt.Println(redisKey)

	exists, err := global.Redis.Exists(ctx, redisKey).Result()
	if err != nil {
		return false
	}
	if exists > 0 {
		return false
	}

	err = global.Redis.Set(ctx, redisKey, 1, RedisNonceTTL).Err()
	if err != nil {
		return false
	}

	return true
}

func HMACMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		signature := c.GetHeader("X-Signature")
		timestamp := c.GetHeader("X-Timestamp")
		nonce := c.GetHeader("X-Nonce")

		if signature == "" || timestamp == "" || nonce == "" {
			response.ErrorBadRequestResponse(c, 4000, "Missing required headers (X-Signature, X-Timestamp, X-Nonce)")
			c.Abort()
			return
		}

		reqTime, err := time.Parse(time.RFC3339, timestamp)
		if err != nil || time.Since(reqTime) > NonceTTL {
			response.ErrorBadRequestResponse(c, 4000, "Invalid or expired X-Timestamp header")
			c.Abort()
			return
		}

		if !ValidateNonce(nonce) {
			response.ErrorBadRequestResponse(c, 4000, "Invalid or expired X-Nonce header")
			c.Abort()
			return
		}

		expectedSignature := GenerateHMACSignature(timestamp + nonce)

		if signature != expectedSignature {
			response.ErrorBadRequestResponse(c, 4000, "Invalid X-Signature header")
			c.Abort()
			return
		}

		c.Next()
	}
}
