package komikcast

import (
	"bacakomik/tools/provider"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync/atomic"

	"github.com/PuerkitoBio/goquery"
	"github.com/iain17/go-cfscrape"
	"github.com/rs/zerolog/log"
)

func GetDates(input string) []string {
	// Define the regular expression pattern for dates
	pattern := `([A-Za-z]{3}\s\d{1,2},\s\d{4})`

	// Compile the regular expression pattern
	reg := regexp.MustCompile(pattern)

	// Find all occurrences of the pattern in the input string
	matches := reg.FindAllString(input, -1)

	return matches
}

// LoadFromFile function
func LoadFromFile() []*provider.Mangalist {
	b, err := os.ReadFile("./dataset/mangalist.json")
	if err != nil {
		log.Err(err).Msg("")
	}

	r := strings.NewReader(string(b))
	var mangas []*provider.Mangalist
	d := json.NewDecoder(r)
	if err := d.Decode(&mangas); err != nil {
		log.Err(err).Msg("Fail decode data from file")
	}
	return mangas
}

// LoadFromLiveUrl function
func LoadFromLiveUrl() []*provider.Mangalist {
	p := provider.NewProvider()
	counter := atomic.Int64{}
	for i := 0; i < 283; i++ {
		s := fmt.Sprintf("https://komikcast.ch/daftar-komik/page/%d", i)
		p.AddProvider(s)
	}

	komikUrls := make(chan provider.Mangalist)
	done := make(chan struct{})

	go WaitAll(p, done, komikUrls)

	for _, u := range p.SiteURL {
		go visit(u, komikUrls, done, &counter)
	}

	var mangalist []*provider.Mangalist
	for m := range komikUrls {
		mangalist = append(mangalist, &provider.Mangalist{
			Title:    m.Title,
			Url:      m.Url,
			ImageUrl: m.ImageUrl,
		})
	}
	provider.SaveJSON("./dataset/mangalist.json", mangalist)

	fmt.Printf("counter: %+v\n", &counter)
	return mangalist
}

func WaitAll(p *provider.Provider, done chan struct{}, komikUrls chan provider.Mangalist) {
	go func() {
		// Receive signals from the done channel
		for range p.SiteURL {
			<-done
		}
		close(komikUrls)
	}()
}

func visit(url string, komikUrls chan provider.Mangalist, done chan struct{}, couunter *atomic.Int64) {
	defer func() {
		done <- struct{}{}
	}()

	r, err := cfscrape.Get(url)
	if err != nil {
		log.Err(err).Msg("")
	}
	defer r.Body.Close()

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Err(err).Msg("[goquery]: error read from reader")
	}

	doc.Find(".list-update_items-wrapper a").Each(func(_ int, s *goquery.Selection) {
		val, exist := s.Attr("href")
		if !exist || val == "" {
			return
		}
		title := s.Find("a .list-update_item-info h3")
		ret, err := title.Html()
		if ret == "" || err != nil {
			return
		}

		imageUrl, exist := s.Find("img").Attr("src")
		if !exist {
			return
		}

		couunter.Add(1)
		komikUrls <- provider.Mangalist{
			Title:    ret,
			Url:      val,
			ImageUrl: imageUrl,
		}
	})

}
func CreateSlug(input string) string {
	// Remove non-alphanumeric characters
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	slug := reg.ReplaceAllString(input, " ")

	// Convert to lower case
	slug = strings.ToLower(slug)

	// Replace spaces with dashes
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove leading and trailing dashes
	slug = strings.Trim(slug, "-")

	return slug
}
