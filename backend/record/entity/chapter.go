package entity

// Chapter struct
type Chapter struct {
	ID      int    `json:"id,omitempty" db:"id"`
	MangaID int    `json:"manga_id,omitempty" db:"mangas_id"`
	Chapter string `json:"chapter,omitempty" db:"chapter"`
	Content string `json:"content,omitempty" db:"content"`
}

type ChapterWithMedia struct {
	Chapter Chapter `json:"chapter,omitempty"`
	Medias  []Media `json:"medias,omitempty"`
}

func (ch *Chapter) String() string {
	return "chapters"
}
