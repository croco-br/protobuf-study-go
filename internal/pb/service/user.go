package service

import (
	"context"

	"github.com/croco-br/protobuf-study-go/internal/pb"
	"github.com/croco-br/protobuf-study-go/internal/pb/database"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	UserDB database.User
}

func NewUserService(userDB database.User) *UserService {
	return &UserService{
		UserDB: userDB,
	}
}

func (u *UserService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user, err := u.UserDB.Create(in.Name, in.Email)
	if err != nil {
		return nil, err
	}

	userResponse := &pb.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: *user.Email,
	}

	return &pb.UserResponse{User: userResponse}, nil
}
