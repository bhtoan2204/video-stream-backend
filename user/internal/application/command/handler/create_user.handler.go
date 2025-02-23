package handler

import (
	"time"

	"github.com/bhtoan2204/user/global"
	"github.com/bhtoan2204/user/internal/application/command/command"
	"github.com/bhtoan2204/user/internal/application/interfaces"
	"github.com/bhtoan2204/user/internal/domain/event"
	infraEvent "github.com/bhtoan2204/user/internal/infrastructure/event"
	"go.uber.org/zap"
)

type CreateUserCommandHandler struct {
	userService    interfaces.UserServiceInterface
	eventPublisher event.DomainEventPublisher
}

func NewCreateUserCommandHandler(userService interfaces.UserServiceInterface) *CreateUserCommandHandler {
	return &CreateUserCommandHandler{
		userService:    userService,
		eventPublisher: infraEvent.NewKafkaDomainEventPublisher(),
	}
}

func (h *CreateUserCommandHandler) Handle(cmd *command.CreateUserCommand) (*command.CreateUserCommandResult, error) {
	result, err := h.userService.CreateUser(cmd)

	userCreatedEvent := event.UserCreatedEvent{
		Payload:  result.Result,
		Occurred: time.Now(),
	}

	go func() {
		if err := h.eventPublisher.Publish(userCreatedEvent); err != nil {
			global.Logger.Error("Failed to publish user created event", zap.Error(err))
			// TODO: Handle error or rollback here
		}
	}()

	return result, err
}
