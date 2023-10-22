package provider

import (
	"fmt"
	"regexp"
	"strings"
	"sync/atomic"

	"github.com/PuerkitoBio/goquery"
	"github.com/iain17/go-cfscrape"
	"github.com/rs/zerolog/log"
)

func Start() {
	p := NewProvider()

	counter := atomic.Int64{}

	for i := 0; i < 1; i++ {
		s := fmt.Sprintf("https://komikcast.ch/daftar-komik/page/%d", i)
		p.AddProvider(s)
	}

	komikUrls := make(chan Mangalist)
	done := make(chan struct{})

	go WaitAll(p, done, komikUrls)

	for _, u := range p.SiteURL {
		go visit(u, komikUrls, done, &counter)
	}

	var mangalist []*Mangalist
	for m := range komikUrls {
		mangalist = append(mangalist, &Mangalist{
			Title:    m.Title,
			Url:      m.Url,
			ImageUrl: m.ImageUrl,
		})
	}
	saveJSON("./dataset/mangalist.json", mangalist)

	fmt.Printf("counter: %+v\n", &counter)

}

func WaitAll(p *Provider, done chan struct{}, komikUrls chan Mangalist) {
	go func() {
		// Receive signals from the done channel
		for range p.SiteURL {
			<-done
		}
		close(komikUrls)
	}()
}

func visit(url string, komikUrls chan Mangalist, done chan struct{}, couunter *atomic.Int64) {
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
		komikUrls <- Mangalist{
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
