package model

type CreateMangaRequest struct {
	Title        string `json:"title,omitempty"`
	Status       string `json:"status,omitempty"`
	ReleaseDate  string `json:"release_date,omitempty"`
	TotalChapter int    `json:"total_chapter,omitempty"`
	Author       string `json:"author,omitempty"`
	Type         string `json:"type,omitempty"`
	Sinopsis     string `json:"sinopsis,omitempty"`
}
