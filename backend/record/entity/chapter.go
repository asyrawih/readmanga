package entity

// Chapter struct
type Chapter struct {
	ID      int    `json:"id,omitempty"`
	MangaID int    `json:"manga_id,omitempty"`
	Chapter string `json:"chapter,omitempty"`
	Content string `json:"content,omitempty"`
}

type ChapterWithMedia struct {
	Chapter Chapter `json:"chapter,omitempty"`
	Medias  []Media `json:"medias,omitempty"`
}

func (ch *Chapter) String() string {
	return "chapters"
}