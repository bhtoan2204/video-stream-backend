package server

import "github.com/bhtoan2204/gateway/pkg/grpc/proto/user"

type Server struct {
	userpb user.UserServiceServer
}

func NewServer(userpb user.UserServiceServer) *Server {
	return &Server{userpb: userpb}
}
