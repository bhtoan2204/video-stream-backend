package middleware

import (
	"context"
	"strings"
	"time"

	"github.com/bhtoan2204/gateway/global"
	"github.com/bhtoan2204/gateway/pkg/grpc/proto/user"
	"github.com/bhtoan2204/gateway/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		parts := strings.Split(c.GetHeader("Authorization"), "Bearer ")
		if (len(parts) < 2) || (parts[1] == "") {
			response.ErrorUnauthorizedResponse(c, response.ErrorUnauthorized)
			c.Abort()
			return
		}
		accessToken := parts[1]
		if global.UserGRPCClient == nil {
			response.ErrorInternalServerResponse(c, response.ErrorInternalServer)
			c.Abort()
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		req := &user.ValidateUserRequest{JwtToken: accessToken}
		user, err := global.UserGRPCClient.ValidateUser(ctx, req)
		if err != nil {
			response.ErrorUnauthorizedResponse(c, response.ErrorUnauthorized)
			c.Abort()
			return
		}

		c.Request.Header.Set("X-User-ID", user.GetId())

		c.Next()
	}
}
