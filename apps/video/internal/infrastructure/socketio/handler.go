package socketio

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	server *Server
}

func NewHandler(server *Server) *Handler {
	return &Handler{
		server: server,
	}
}

// ServeHTTP implements http.Handler interface
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.server.server.ServeHTTP(w, r)
}

// HandleSocketIO is a convenience method for Gin handlers
func (h *Handler) HandleSocketIO(c *gin.Context) {
	h.ServeHTTP(c.Writer, c.Request)
}
