package domain_service

import (
	repository "github.com/bhtoan2204/user/internal/domain/repository/command"
	"github.com/bhtoan2204/user/utils"
)

type AuthDomainService struct {
	userRepo repository.UserRepository
}

func NewAuthDomainService(userRepo repository.UserRepository) AuthDomainService {
	return AuthDomainService{
		userRepo: userRepo,
	}
}

func (s AuthDomainService) IsUsernameTaken(username string) (bool, error) {
	user, err := s.userRepo.FindOneByQuery(utils.QueryOptions{
		Filters: map[string]interface{}{
			"username": username,
		},
	})
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
