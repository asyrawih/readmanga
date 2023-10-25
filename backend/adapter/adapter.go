package adapter

import (
	"context"
)

type Accessable[T any] interface {
	// Get Access to instance of of T
	NewApi() *T
}

type Creational[T any, K any] interface {
	// Create Data
	Create(ctx context.Context, data *T) (K, error)
}

type Modificational[T any, K any] interface {
	// Update Data
	Update(ctx context.Context, data *T, id K) error
}

type Retrival[T any, K any] interface {
	// Get All Data
	GetAll(ctx context.Context) []*T
	// Retrive One Data
	GetOne(ctx context.Context, id K) *T
}

type Destroyer[ID any] interface {
	// Delete the record
	Delete(ctx context.Context, id ID) bool
}
