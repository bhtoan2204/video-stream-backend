package listener

import (
	"context"

	"github.com/bhtoan2204/user/global"
	common "github.com/bhtoan2204/user/internal/application/common/event"
	"github.com/bhtoan2204/user/internal/application/event_bus/event"
	"github.com/bhtoan2204/user/internal/domain/entities"
	repository "github.com/bhtoan2204/user/internal/domain/repository/command"
	eSRepository "github.com/bhtoan2204/user/internal/domain/repository/query"
)

type UserListener struct {
	userRepository   repository.UserRepositoryInterface
	esUserRepository eSRepository.ESUserRepositoryInterface
}

func NewUserListener(userRepository repository.UserRepositoryInterface, esUserRepository eSRepository.ESUserRepositoryInterface) *UserListener {
	return &UserListener{
		userRepository:   userRepository,
		esUserRepository: esUserRepository,
	}
}

func (l *UserListener) IndexUser(ctx context.Context, indexUserEvent *event.IndexUserEvent) (*common.IndexResult, error) {
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

	if err := l.esUserRepository.Index(ctx, user); err != nil {
		return nil, err
	}
	global.Logger.Info("User indexed successfully")
	return &common.IndexResult{
		Success: true,
	}, nil
}
