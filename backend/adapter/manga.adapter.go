package adapter

import (
	"bacakomik/record/entity"
	"bacakomik/repository/mysql"
)

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
