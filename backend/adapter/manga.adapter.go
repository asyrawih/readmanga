package adapter

import (
	"bacakomik/record/entity"
	"bacakomik/record/model"
	"bacakomik/repository/mysql"
	"context"
)

type GetMangaChapters interface {
	GetMangaWithChapters(ctx context.Context, id int) model.GetMangaChapter
}

// Contract of Service manga  as depedency
type ServiceMangaCreational interface {
	Creational[entity.Manga, int]
	Modificational[entity.Manga, int]
	Retrival[entity.Manga, int]
	Destroyer[int]
	GetMangaChapters
}

// Contract MAnga Repo
type RepoMangaCreational interface {
	Creational[entity.Manga, int]
	Modificational[entity.Manga, int]
	Retrival[entity.Manga, int]
	Destroyer[int]
	Accessable[mysql.MangaRepository]
}
