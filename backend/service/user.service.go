package service

import (
	"context"

	"bacakomik/adapter"
	"bacakomik/record/entity"
)

type UserService struct {
	repo adapter.RepoUserCreational
}

// NewUserService function
func NewUserService(repo adapter.RepoUserCreational) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Create Data
func (us *UserService) Create(ctx context.Context, data *entity.User) error {
	err := us.repo.Create(ctx, data)
	return err
}
