package listener

import (
	"github.com/bhtoan2204/user/global"
	common "github.com/bhtoan2204/user/internal/application/common/event"
	"github.com/bhtoan2204/user/internal/application/event/event"
	"github.com/bhtoan2204/user/internal/domain/entities"
	repository "github.com/bhtoan2204/user/internal/domain/repository/command"
	eSRepository "github.com/bhtoan2204/user/internal/domain/repository/query"
)

type UserListener struct {
	userRepository   repository.UserRepository
	esUserRepository eSRepository.ESUserRepository
}

func NewUserListener(userRepository repository.UserRepository, esUserRepository eSRepository.ESUserRepository) *UserListener {
	return &UserListener{
		userRepository:   userRepository,
		esUserRepository: esUserRepository,
	}
}

func (l *UserListener) IndexUser(indexUserEvent *event.IndexUserEvent) (*common.IndexResult, error) {
	user := &entities.User{
		AbstractModel: entities.AbstractModel{
			ID:        indexUserEvent.ID,
			CreatedAt: indexUserEvent.CreatedAt,
			UpdatedAt: indexUserEvent.UpdatedAt,
			DeletedAt: indexUserEvent.DeletedAt,
		},
		Username:     indexUserEvent.Username,
		Email:        indexUserEvent.Email,
		FirstName:    indexUserEvent.FirstName,
		LastName:     indexUserEvent.LastName,
		Phone:        indexUserEvent.Phone,
		BirthDate:    indexUserEvent.BirthDate,
		Address:      indexUserEvent.Address,
		PasswordHash: indexUserEvent.PasswordHash,
		PinCode:      indexUserEvent.PinCode,
		Status:       1,
	}

	if err := l.esUserRepository.Index(user); err != nil {
		return nil, err
	}
	global.Logger.Info("User indexed successfully")
	return &common.IndexResult{
		Success: true,
	}, nil
}
