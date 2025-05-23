package handler

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/bhtoan2204/gateway/internal/consul"
	"github.com/bhtoan2204/gateway/internal/modules/auth/dto"
	"github.com/bhtoan2204/gateway/pkg/response"
	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary      Login
// @Description  Login to the system
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "Login Request"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ResponseData
// @Router       /user-service/auth/login [post]
func Login(c *gin.Context) {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)

	var req dto.LoginRequest
	if err := json.NewDecoder(tee).Decode(&req); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}
	if err := req.Validate(); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(buf.Bytes()))
	consul.ServiceProxy("user-service")(c)
}

// RefreshToken godoc
// @Summary      Refresh Token
// @Description  Refresh Token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dto.RefreshTokenRequest true "Refresh Token Request"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ResponseData
// @Router       /user-service/auth/refresh [post]
func RefreshToken(c *gin.Context) {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)

	var req dto.RefreshTokenRequest
	if err := json.NewDecoder(tee).Decode(&req); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}

	if err := req.Validate(); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(buf.Bytes()))
	consul.ServiceProxy("user-service")(c)
}

// Logout godoc
// @Summary      Logout
// @Description  Logout from the system
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LogoutRequest true "Logout Request"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ResponseData
// @Router       /user-service/auth/logout [post]
func Logout(c *gin.Context) {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)

	var req dto.LogoutRequest
	if err := json.NewDecoder(tee).Decode(&req); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}

	if err := req.Validate(); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(buf.Bytes()))
	consul.ServiceProxy("user-service")(c)
}

// Setup2FA godoc
// @Summary      Setup 2FA
// @Description  Setup 2FA
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ResponseData
// @Router       /user-service/auth/2fa/setup [post]
func Setup2FA(c *gin.Context) {
	consul.ServiceProxy("user-service")(c)
}

// Verify2FA godoc
// @Summary      Verify 2FA
// @Description  Verify 2FA
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ResponseData
// @Router       /user-service/auth/2fa/verify [post]
func Verify2FA(c *gin.Context) {
	consul.ServiceProxy("user-service")(c)
}
