package service

import (
	"context"
	"testing"

	mocks "bacakomik/mocks/adapter"
	"bacakomik/record/entity"
)

func TestNewUserService(t *testing.T) {
	// ruc := mocks.NewServiceMangaCreational(t)
	// us := NewUserService(ruc)
	// assert.NotNil(t, us)
}

func setupUserTest(t *testing.T, ruc *mocks.RepoUserCreational) (*UserService, *entity.User, context.Context) {
	user := &entity.User{
		Name:     "hanan",
		Email:    "hasyrawi@gmail.com",
		Username: "hanan",
		Password: "randomstring",
	}
	ctx := context.Background()
	return &UserService{repo: ruc}, user, ctx
}

func TestUserService_Create(t *testing.T) {
	tests := []struct {
		name      string
		setup     func(ruc *mocks.RepoUserCreational) (*UserService, *entity.User, context.Context)
		wantError bool
	}{
		{
			name: "should success create user",
			setup: func(ruc *mocks.RepoUserCreational) (*UserService, *entity.User, context.Context) {
				us, u, ctx := setupUserTest(t, ruc)
				ruc.On("Create", ctx, u).Return(99, nil)
				return us, u, ctx
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ruc := mocks.NewRepoUserCreational(t)
			us, u, ctx := tt.setup(ruc)
			_, err := us.Create(ctx, u)
			if (err != nil) != tt.wantError {
				t.Errorf("expected %v, actual %v", tt.wantError, err)
			}
		})
	}
}
