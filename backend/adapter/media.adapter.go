package adapter

import (
	"bacakomik/record/entity"
	"bacakomik/repository/mysql"
)

// Contract of Service manga  as depedency
type ServiceMediaCreational interface {
	Creational[entity.Media, int]
	Modificational[entity.Media, int]
	Retrival[entity.Media, int]
	Destroyer[int]
}

// Contract MAnga Repo
type RepoMediaCreational interface {
	Creational[entity.Media, int]
	Modificational[entity.Media, int]
	Retrival[entity.Manga, int]
	Destroyer[int]
	Accessable[mysql.MediaRepository]
}
