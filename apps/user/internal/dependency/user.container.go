package dependency

import (
	"errors"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/domain/interfaces"
	writeRepository "github.com/bhtoan2204/user/internal/domain/repository/command"
	readRepository "github.com/bhtoan2204/user/internal/domain/repository/query"
	eSRepository "github.com/bhtoan2204/user/internal/infrastructure/db/elasticsearch/repository"
	"github.com/bhtoan2204/user/internal/infrastructure/db/mysql/repository"
	"github.com/bhtoan2204/user/internal/infrastructure/listener"
)

type UserContainer struct {
	UserRepository   writeRepository.UserRepositoryInterface
	ESUserRepository readRepository.ESUserRepositoryInterface
	UserListener     interfaces.UserListenerInterface
}

func BuildUserContainer() (*UserContainer, error) {
	if global.MDB == nil {
		return nil, errors.New("Write repository is required")
	}
	if global.ESClient == nil {
		return nil, errors.New("Read repository is required")
	}

	eSUserRepository := eSRepository.NewESUserRepository(global.ESClient)
	userRepository := repository.NewUserRepository(global.MDB)
	userListener := listener.NewUserListener(userRepository, eSUserRepository)

	return &UserContainer{
		UserRepository:   userRepository,
		ESUserRepository: eSUserRepository,
		UserListener:     userListener,
	}, nil
}
