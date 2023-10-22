package provider

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

// Provider struct
// provider will contain the website want to scraper
// with some configuration
type Provider struct {
	SiteURL []string
}

// NewProvider function
func NewProvider() *Provider {
	return &Provider{}
}

// AddProvider method
// Add New Url into prpovider
func (p *Provider) AddProvider(url string) {
	p.SiteURL = append(p.SiteURL, url)
}

// saveJSON function
func saveJSON(filename string, data []*Mangalist) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Err(err).Msg("Error encoding JSON")
		return
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Err(err).Msg("Error writing JSON file")
		return
	}

	fmt.Println("Data saved to", filename)
}
