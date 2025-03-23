package handler

import (
	"github.com/bhtoan2204/gateway/internal/consul"
	"github.com/bhtoan2204/gateway/internal/modules/user/dto"
	"github.com/bhtoan2204/gateway/pkg/response"
	"github.com/gin-gonic/gin"
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
	var req dto.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}

	if err := req.Validate(); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}

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
	var req dto.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}

	if err := req.Validate(); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}

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
	var req dto.SearchUserRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}

	if err := req.Validate(); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}

	req.SetDefaults()

	consul.ServiceProxy("user-service")(c)
}
