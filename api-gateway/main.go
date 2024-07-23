package main

import (
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	pb "github.com/mashinolol/auth-service/service/proto"
	"github.com/mashinolol/auth-service/service/user"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	db, err := sqlx.Connect("postgres", "user=postgres password=mysecretpassword dbname=mydb sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	userService := user.NewUserService(db)
	userHandler := user.NewUserHandler(userService)

	pb.RegisterUserServiceServer(s, userHandler)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
