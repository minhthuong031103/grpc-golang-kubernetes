package internal

import (
	"context"
	users "user-service/protogen/golang/user"
)

type UserService struct {
	users.UnimplementedUserServiceServer
	db *DB
}

func NewUserService(db *DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetUser(ctx context.Context, req *users.GetUserRequest) (*users.GetUserResponse, error) {
	// Example logic to fetch user from db using req.UserId
	user, err := s.db.GetUser(req.UserId)
	if err != nil {
		return nil, err
	}

	return &users.GetUserResponse{
		User: &users.User{
			UserId: user.ID,
			Name:   user.Name,
			Email:  user.Email,
		},
	}, nil
}

// Implement other methods if needed...
