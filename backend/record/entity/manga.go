package entity

import (
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type Manga struct {
	Title       string `json:"title,omitempty"`
	Author      string `json:"author,omitempty"`
	ReleaseDate string `json:"release_date,omitempty"`
	Rating      int    `json:"rating,omitempty"`
}

// Convert The Struct into json
func (m *Manga) ToJSon() string {
	b, err := json.Marshal(m)
	if err != nil {
		log.Printf("error: %s", err.Error())
	}
	return string(b)
}

func (m *Manga) ToBson() string {
	panic("implement me")
}
