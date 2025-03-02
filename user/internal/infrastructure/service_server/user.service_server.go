package service_server

import (
	"context"
	"fmt"

	"github.com/bhtoan2204/user/global"
	repository "github.com/bhtoan2204/user/internal/domain/repository/command"
	gprcUser "github.com/bhtoan2204/user/internal/infrastructure/grpc/proto/user"
	"github.com/bhtoan2204/user/pkg/jwt_utils"
	"github.com/bhtoan2204/user/utils"
	"go.uber.org/zap"
)

type UserServiceServerImpl struct {
	gprcUser.UnimplementedUserServiceServer
	userRepository repository.UserRepository
}

func NewUserServiceServer(userRepository repository.UserRepository) gprcUser.UserServiceServer {
	return &UserServiceServerImpl{
		userRepository: userRepository,
	}
}

func (s *UserServiceServerImpl) ValidateUser(ctx context.Context, req *gprcUser.ValidateUserRequest) (*gprcUser.UserResponse, error) {
	if req.JwtToken == "" {
		return nil, fmt.Errorf("jwt token empty")
	}

	parsedToken, err := jwt_utils.ExtractTokenClaims(req.JwtToken, global.Config.SecurityConfig.JWTAccessSecret)

	if err != nil {
		global.Logger.Error("Failed to parse jwt token", zap.Error(err))
		return nil, fmt.Errorf("failed to parse jwt token: %v", err)
	}

	userID, ok := parsedToken["id"].(string)

	if !ok {
		global.Logger.Error("user_id not found in jwt token")
		return nil, fmt.Errorf("user_id not found in jwt token")
	}

	user, err := s.userRepository.FindOneByQuery(
		&utils.QueryOptions{
			Filters: map[string]interface{}{"id": userID},
		},
	)

	if err != nil {
		global.Logger.Error("Failed to find user", zap.Error(err))
		return nil, fmt.Errorf("failed to find user: %v", err)
	}

	var roles []*gprcUser.Role
	for _, r := range user.Roles {
		var grpcPermissions []*gprcUser.Permission
		for _, p := range r.Permissions {
			grpcPermissions = append(grpcPermissions, &gprcUser.Permission{
				Id:          p.ID,
				Name:        p.Name,
				Description: p.Description,
			})
		}
		roles = append(roles, &gprcUser.Role{
			Id:          r.ID,
			Name:        r.Name,
			Permissions: grpcPermissions,
		})
	}

	return &gprcUser.UserResponse{
		Id:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Roles:     roles,
	}, nil
}

func (s *UserServiceServerImpl) mustEmbedUnimplementedUserServiceServer() {}
