package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	"github.com/bhtoan2204/gateway/global"
	"github.com/bhtoan2204/gateway/internal/consul"
	"github.com/bhtoan2204/gateway/internal/modules/user/dto"
	"github.com/bhtoan2204/gateway/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateUser godoc
// @Summary      Create a new user
// @Description  Create a new user with the given details
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.CreateUserRequest  true  "User details"
// @Success      201   {object}  response.ResponseData
// @Failure      400   {object}  response.ResponseData
// @Failure      500   {object}  response.ResponseData
// @Router       /user-service/users [post]
func CreateUser(c *gin.Context) {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)

	var req dto.CreateUserRequest
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

// UpdateUser godoc
// @Summary      Update user details
// @Description  Update user details with the given information
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      dto.UpdateProfileRequest  true  "User details"
// @Success      200   {object}  response.ResponseData
// @Failure      400   {object}  response.ResponseData
// @Failure      500   {object}  response.ResponseData
// @Router       /user-service/users [put]
func UpdateUser(c *gin.Context) {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)

	var req dto.UpdateProfileRequest
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

// SearchUser godoc
// @Summary      Search users
// @Description  Search users with the given query parameters
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        query          query     string  true   "Search query"
// @Param        page           query     int     false  "Page number"
// @Param        limit          query     int     false  "Limit"
// @Param        sort_by        query     string  false  "Sort by"
// @Param        sort_direction query     string  false  "Sort direction"
// @Success      200 {object} response.ResponseData
// @Failure      400 {object} response.ResponseData
// @Failure      500 {object} response.ResponseData
// @Router       /user-service/users/search [get]
func GetUserProfile(c *gin.Context) {
	consul.ServiceProxy("user-service")(c)
}

// SearchUser godoc
// @Summary Search users
// @Description Search users with the given details
// @Tags users
// @Accept json
// @Produce json
// @Param query query string true "Search query"
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Param sort_by query string false "Sort by"
// @Param sort_direction query string false "Sort direction"
// @Success 200 {object} response.ResponseData
// @Failure 400 {object} response.ResponseData
// @Failure 500 {object} response.ResponseData
// @Router /user-service/users/search [get]
func SearchUser(c *gin.Context) {
	// Log query parameters
	queryParams := c.Request.URL.Query()
	for key, values := range queryParams {
		global.Logger.Info("Query parameter",
			zap.String("key", key),
			zap.Strings("values", values))
	}

	// Parse page and limit
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	// Create request from query parameters
	req := dto.SearchUserRequest{
		Query:         c.Query("query"),
		Page:          page,
		Limit:         limit,
		SortBy:        c.Query("sort_by"),
		SortDirection: c.Query("sort_direction"),
	}

	if err := req.Validate(); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}

	req.SetDefaults()

	// Forward request to user service
	consul.ServiceProxy("user-service")(c)
}
