package serviceserver

import (
	"context"
	"fmt"

	"github.com/bhtoan2204/user/internal/infrastructure/grpc/proto/user"
)

type UserServiceServerImpl struct {
	user.UnimplementedUserServiceServer
}

func NewUserServiceServer() user.UserServiceServer {
	return &UserServiceServerImpl{}
}

func (s *UserServiceServerImpl) ValidateUser(ctx context.Context, req *user.ValidateUserRequest) (*user.UserResponse, error) {
	fmt.Println("ValidateUserrrrrrrrrrrrrrrrrr")
	if req.JwtToken == "" {
		return nil, fmt.Errorf("jwt token trống")
	}

	return &user.UserResponse{
		Id:        "1",
		Username:  "mockuser",
		Email:     "mockuser@example.com",
		FirstName: "Mock",
		LastName:  "User",
		Phone:     "0123456789",
		Roles: []*user.Role{
			{
				Id:   "role1",
				Name: "admin",
				Permissions: []*user.Permission{
					{
						Id:          "perm1",
						Name:        "read",
						Description: "Quyền đọc dữ liệu",
					},
					{
						Id:          "perm2",
						Name:        "write",
						Description: "Quyền ghi dữ liệu",
					},
				},
			},
		},
	}, nil
}

func (s *UserServiceServerImpl) mustEmbedUnimplementedUserServiceServer() {}
