package model

type CreateChapterRequest struct {
	MangaID int    `json:"manga_id,omitempty"`
	Chapter string `json:"chapter,omitempty"`
	Content string `json:"content,omitempty"`
}
