package handler

import (
	"bytes"
	"encoding/json"
	"io"

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
// @Success      201  {object}  response.ResponseData
// @Failure      400  {object}  response.ResponseData
// @Failure      401  {object}  response.ResponseData
// @Router       /comment-service/comments [post]
func CreateComment(c *gin.Context) {
	var buf bytes.Buffer
	tee := io.TeeReader(c.Request.Body, &buf)

	var req dto.CreateCommentRequest
	if err := json.NewDecoder(tee).Decode(&req); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}

	if err := req.Validate(); err != nil {
		response.ErrorBadRequestResponse(c, response.ErrorBadRequest, err)
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(buf.Bytes()))
	consul.ServiceProxy("comment-service")(c)
}
