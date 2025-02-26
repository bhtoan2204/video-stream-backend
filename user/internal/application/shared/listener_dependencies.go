package shared

import "github.com/bhtoan2204/user/internal/domain/interfaces"

type ListenerDependencies struct {
	UserListener interfaces.UserListenerInterface
	// Other services can be added here
}
