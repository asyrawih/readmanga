package adapter

import (
	"context"

	"bacakomik/record/entity"
	"bacakomik/repository/mysql"
)

// ChapterRepoCreational interface
type ServiceChapterCreational interface {
	Creational[entity.Chapter, int]
	Modificational[entity.Chapter, int]
	Retrival[entity.Chapter, int]
	Destroyer[int]
}

// ChapterRepoCreational interface
type ChapterRepoCreational interface {
	Creational[entity.Chapter, int]
	Modificational[entity.Chapter, int]
	Retrival[entity.Chapter, int]
	Destroyer[int]
}

// Contract of Service manga  as depedency
type ServiceMangaCreational interface {
	Creational[entity.Manga, int]
	Modificational[entity.Manga, int]
	Retrival[entity.Manga, int]
	Destroyer[int]
}

// Contract MAnga Repo
type RepoMangaCreational interface {
	Creational[entity.Manga, int]
	Modificational[entity.Manga, int]
	Retrival[entity.Manga, int]
	Destroyer[int]
	Accessable[mysql.MangaRepository]
}

// Contract User Repo
type RepoUserCreational interface {
	Creational[entity.User, int]
	Modificational[entity.User, int]
	Retrival[entity.User, int]
	Destroyer[int]
	Accessable[mysql.UserRepository]
}

// Contract User Repo
type ServiceUserCreational interface {
	Creational[entity.User, int]
	Modificational[entity.User, int]
	Retrival[entity.User, int]
	Destroyer[int]
}

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
