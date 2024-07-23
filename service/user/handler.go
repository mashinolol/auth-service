package user

import (
	"context"

	pb "github.com/mashinolol/auth-service/service/proto"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return h.service.CreateUser(ctx, req)
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return h.service.GetUser(ctx, req)
}

func (h *UserHandler) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return h.service.UpdateUser(ctx, req)
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return h.service.DeleteUser(ctx, req)
}
