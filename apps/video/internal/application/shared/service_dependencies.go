package shared

import (
	"github.com/bhtoan2204/video/internal/domain/interfaces"
	"github.com/bhtoan2204/video/internal/infrastructure/socketio"
)

type ServiceDependencies struct {
	VideoService  interfaces.VideoServiceInterface
	SocketService *socketio.Service
	// Other services can be added here
}
