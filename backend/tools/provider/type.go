package provider

type Mangalist struct {
	Title    string `json:"title,omitempty"`
	Url      string `json:"url,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
}
