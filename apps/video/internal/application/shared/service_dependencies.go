package shared

import "github.com/bhtoan2204/video/internal/domain/interfaces"

type ServiceDependencies struct {
	VideoService interfaces.VideoServiceInterface
	// Other services can be added here
}
