package entity

import (
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"
)

type Manga struct {
	ID           int       `json:"id,omitempty"           `
	Title        string    `json:"title,omitempty"        `
	Status       string    `json:"status,omitempty"       `
	ReleaseDate  string    `json:"release_date,omitempty" `
	TotalChapter int       `json:"total_chapter,omitempty"`
	Author       string    `json:"author,omitempty"       `
	Type         string    `json:"type,omitempty"         `
	Sinopsis     string    `json:"sinopsis,omitempty"     `
	CreatedBy    int       `json:"created_by,omitempty"   `
	CreatedAt    time.Time `json:"created_at,omitempty"   `
}

// Convert The Struct into json
func (m *Manga) ToJSon() string {
	b, err := json.Marshal(m)
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
	return string(b)
}
