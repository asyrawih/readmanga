package entity

type Media struct {
	ID        int    `json:"id,omitempty"`
	ModelType string `json:"model_type,omitempty"`
	ModelID   int    `json:"model_id,omitempty"`
	URL       string `json:"url,omitempty"`
}

func (me *Media) String() string {
	return "media"
}
