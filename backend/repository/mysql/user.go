package mysql

import (
	"context"

	"github.com/jackc/pgx/v5"

	"bacakomik/record/entity"
)

// UserRepository struct
type UserRepository struct {
	conn *pgx.Conn
}

// NewUserRepository function
func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

// Create Data
func (us *UserRepository) Create(ctx context.Context, data *entity.User) error {
	return nil
}

// Get All Data
func (us *UserRepository) GetAll(ctx context.Context) []*entity.User {
	panic("not implemented") // TODO: Implement
}

// Retrive One Data
func (us *UserRepository) GetOne(ctx context.Context, id int) *entity.User {
	panic("not implemented") // TODO: Implement
}

func (us *UserRepository) Delete(ctx context.Context, id int) bool {
	panic("not implemented") // TODO: Implement
}

// Update Data
func (us *UserRepository) Update(ctx context.Context, data *entity.User, id int) error {
	panic("not implemented") // TODO: Implement
}

// Get Access to instance of of T
func (us *UserRepository) NewApi() *UserRepository {
	return us
}
