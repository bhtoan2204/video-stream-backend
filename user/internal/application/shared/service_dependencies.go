package shared

import "github.com/bhtoan2204/user/internal/domain/interfaces"

type ServiceDependencies struct {
	UserService interfaces.UserServiceInterface
	// Other services can be added here
}
