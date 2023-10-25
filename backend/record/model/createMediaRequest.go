package model

import (
	"bacakomik/record/entity"
)

// CreateMediaRequest struct
type CreateMediaRequest struct {
	ModelType string `json:"model_type,omitempty" form:"model_type"`
	ModelID   int    `json:"model_id,omitempty" form:"model_id"`
	URL       string `json:"url,omitempty" form:"url"`
}

func (c *CreateMediaRequest) Into() (m entity.Media) {
	return entity.Media{
		ModelType: m.String(),
		ModelID:   m.ModelID,
		URL:       m.URL,
	}
}
