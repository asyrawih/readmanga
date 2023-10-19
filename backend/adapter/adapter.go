package adapter

import (
	"context"

	"bacakomik/record/entity"
)

type ServiceMangaCreational interface {
	Creational[entity.Manga]
	Modificational[entity.Manga, int]
	Retrival[entity.Manga, int]
}

type Creational[T any] interface {
	// Create Data
	Create(ctx context.Context, data *T) error
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
