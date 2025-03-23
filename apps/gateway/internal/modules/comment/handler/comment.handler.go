package handler

import (
	"github.com/bhtoan2204/gateway/internal/consul"
	"github.com/bhtoan2204/gateway/internal/modules/comment/dto"
	"github.com/bhtoan2204/gateway/pkg/response"
	"github.com/gin-gonic/gin"
)

// CreateComment godoc
// @Summary      Create Comment
// @Description  Create a new comment on a video
// @Tags         comments
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateCommentRequest true "Create Comment Request"
// @Success      201  {object}  response.SuccessResponse
// @Failure      400  {object}  response.ResponseData
// @Failure      401  {object}  response.ResponseData
// @Router       /comment-service/comments [post]
func CreateComment(c *gin.Context) {
	var req dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}

	if err := req.Validate(); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}

	// Proxy call to comment-service
	consul.ServiceProxy("comment-service")(c)
}
