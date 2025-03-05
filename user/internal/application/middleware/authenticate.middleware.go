package middleware

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/infrastructure/grpc/proto/user"
	"github.com/bhtoan2204/user/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		encodedUser := c.GetHeader("X-User-Data")
		if encodedUser == "" {
			global.Logger.Error("X-User-Data header is missing")
			response.ErrorUnauthorizedResponse(c, http.StatusUnauthorized)
			c.Abort()
			return
		}
		userData, err := base64.StdEncoding.DecodeString(encodedUser)
		if err != nil {
			global.Logger.Error("Failed to decode user data: ", zap.Error(err))
			response.ErrorUnauthorizedResponse(c, http.StatusUnauthorized)
			c.Abort()
			return
		}
		var userObj user.UserResponse
		if err := json.Unmarshal(userData, &userObj); err != nil {
			global.Logger.Error("Failed to unmarshal user data", zap.Error(err))
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		if userObj.Id == "" {
			global.Logger.Error("User ID is missing")
			response.ErrorUnauthorizedResponse(c, 401)
			return
		}
		c.Set("user", userData)
		c.Next()
	}
}
