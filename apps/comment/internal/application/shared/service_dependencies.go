package shared

import (
	"github.com/bhtoan2204/comment/internal/domain/interfaces"
)

type ServiceDependencies struct {
	CommentService interfaces.CommentServiceInterface
	// Other services can be added here
}
