package adapter

import (
	"bacakomik/record/entity"
	"bacakomik/repository/mysql"
)

// ChapterRepoCreational interface
type ServiceChapterCreational interface {
	Creational[entity.Chapter, int]
	Modificational[entity.Chapter, int]
	Retrival[entity.ChapterWithMedia, int]
	Destroyer[int]
}

// ChapterRepoCreational interface
type ChapterRepoCreational interface {
	Creational[entity.Chapter, int]
	Modificational[entity.Chapter, int]
	Retrival[entity.Chapter, int]
	Destroyer[int]
	Accessable[mysql.ChapterRepositry]
}
