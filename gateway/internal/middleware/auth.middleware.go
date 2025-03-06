package middleware

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/bhtoan2204/gateway/global"
	"github.com/bhtoan2204/gateway/internal/redis"
	"github.com/bhtoan2204/gateway/pkg/grpc/proto/user"
	"github.com/bhtoan2204/gateway/pkg/response"

	"github.com/gin-gonic/gin"
)

const cacheExpiration = 15 * time.Minute // adjust TTL as needed

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the access token from the "Authorization" header
		parts := strings.Split(c.GetHeader("Authorization"), "Bearer ")
		if len(parts) < 2 || parts[1] == "" {
			response.ErrorUnauthorizedResponse(c, response.ErrorUnauthorized)
			c.Abort()
			return
		}
		accessToken := parts[1]
		cacheKey := "auth:user:" + accessToken
		ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
		defer cancel()

		var validatedUser *user.UserResponse

		// Attempt to retrieve user data from Redis
		cachedData, err := redis.GetData(ctx, cacheKey)
		if err == nil && cachedData != "" {
			var cachedUser user.UserResponse
			if err := json.Unmarshal([]byte(cachedData), &cachedUser); err == nil {
				validatedUser = &cachedUser
			}
		}

		if validatedUser == nil {
			if global.UserGRPCClient == nil {
				response.ErrorInternalServerResponse(c, response.ErrorInternalServer)
				c.Abort()
				return
			}

			grpcCtx, grpcCancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
			defer grpcCancel()

			req := &user.ValidateUserRequest{JwtToken: accessToken}
			usr, err := global.UserGRPCClient.ValidateUser(grpcCtx, req)
			if err != nil {
				response.ErrorUnauthorizedResponse(c, response.ErrorUnauthorized)
				c.Abort()
				return
			}
			validatedUser = usr

			// Cache the validated user data
			if userData, err := json.Marshal(validatedUser); err == nil {
				redis.CacheData(ctx, cacheKey, string(userData), int(cacheExpiration.Seconds()))
			}
		}

		userData, err := json.Marshal(validatedUser)
		if err != nil {
			response.ErrorInternalServerResponse(c, response.ErrorInternalServer)
			c.Abort()
			return
		}
		encodedUser := base64.StdEncoding.EncodeToString(userData)
		c.Request.Header.Set("X-User-Data", encodedUser)
		c.Next()
	}
}
