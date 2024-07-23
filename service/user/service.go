package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	pb "github.com/mashinolol/auth-service/service/proto"
)

type UserService struct {
	db *sqlx.DB
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// Implement CreateUser logic
	return &pb.CreateUserResponse{}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// Implement GetUser logic
	return &pb.GetUserResponse{}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	// Implement UpdateUser logic
	return &pb.UpdateUserResponse{}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	// Implement DeleteUser logic
	return &pb.DeleteUserResponse{}, nil
}
