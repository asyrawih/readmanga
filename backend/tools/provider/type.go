package provider

type Mangalist struct {
	Title    string `json:"title,omitempty"`
	Url      string `json:"url,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}

type MangaDetail struct {
	Title       string    `json:"title,omitempty"`
	Status      string    `json:"status,omitempty"`
	ReleaseDate string    `json:"release_date,omitempty"`
	Author      string    `json:"author,omitempty"`
	Type        string    `json:"type,omitempty"`
	Sinopsis    string    `json:"sinopsis,omitempty"`
	Chapter     []Chapter `json:"chapter,omitempty"`
}

type Chapter struct {
	ChapterURl string `json:"chapter_u_rl,omitempty"`
	Chapter    string `json:"chapter,omitempty"`
}
