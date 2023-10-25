package entity

import (
	"encoding/json"
	"time"

	"github.com/rs/zerolog/log"
)

type Manga struct {
	ID           int       `json:"id,omitempty" db:"id"`
	Title        string    `json:"title,omitempty" db:"title"`
	Status       string    `json:"status,omitempty" db:"status"`
	ReleaseDate  string    `json:"release_date,omitempty" db:"release_date"`
	TotalChapter int       `json:"total_chapter,omitempty" db:"total_chapter"`
	Author       string    `json:"author,omitempty" db:"author"`
	Type         string    `json:"type,omitempty" db:"type"`
	Sinopsis     string    `json:"sinopsis,omitempty" db:"sinopsis"`
	CreatedBy    int       `json:"created_by,omitempty" db:"created_by"`
	CreatedAt    time.Time `json:"created_at,omitempty" db:"created_at"`
}

// Convert The Struct into json
func (m *Manga) ToJSon() string {
	b, err := json.Marshal(m)
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
	return string(b)
}

// String method
func (ma *Manga) String() string {
	return "manga"
}
