package model

import "bacakomik/record/entity"

type GetMangaChapter struct {
	Manga    entity.Manga
	Chapters []entity.Chapter
}
