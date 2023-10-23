package komikcast

import (
	"bacakomik/tools/provider"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/iain17/go-cfscrape"
	"github.com/rs/zerolog/log"
)

func ProcessMangaDetail() {
	mangas := LoadFromFile()

	mangaDetail := make(chan *provider.MangaDetail)
	done := make(chan struct{})

	for id, m := range mangas {
		log.Info().Msgf("process %d", id)
		go GetDetail(m, mangaDetail, done)
	}

	go WaitDetail(len(mangas), done, mangaDetail)

	var mangaDetails []*provider.MangaDetail

	for md := range mangaDetail {
		mangaDetails = append(mangaDetails, md)
	}

	provider.SaveJSON("./dataset/manga-detail.json", mangaDetails)

}

func WaitDetail(counter int, done chan struct{}, d chan *provider.MangaDetail) {
	go func() {
		for i := 0; i < counter; i++ {
			<-done
			log.Info().Msgf("done %d", i)
		}
		close(d)
	}()
}

func cleanUrl(rawURl string) string {
	s, err := url.PathUnescape(rawURl)
	if err != nil {
		return rawURl
	}
	return s
}

func GetDetail(mangas *provider.Mangalist, out chan *provider.MangaDetail, done chan struct{}) {
	defer func() {
		done <- struct{}{}
		if r := recover(); r != nil {
			return
		}
	}()
	log.Info().Msgf("fetch: %s", mangas.Url)
	r, err := cfscrape.Get(cleanUrl(mangas.Url))
	if err != nil {
		log.Err(err).Msg("")
	}
	defer r.Body.Close()

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Err(err).Msg("[goquery]: error read from reader")
	}

	doc.Find(".wrapper").Each(func(_ int, s *goquery.Selection) {
		// Get Title
		title, err := s.Find(".komik_info-content-body .komik_info-content-body-title").First().Html()
		if err != nil {
			return
		}
		if title == "" {
			return
		}

		// Get Release Date
		releaseDateNative := s.Find(".komik_info-content-meta .komik_info-content-info-release").Text()
		if releaseDateNative == "" {
			return
		}
		text := GetDates(releaseDateNative)
		release := strings.Join(text, " ")

		// Get Type
		typeHtml, _ := s.Find(".komik_info-content-info-type a").Html()

		var chapter []provider.Chapter
		s.Find(".komik_info-chapters-wrapper li").Each(func(_ int, s *goquery.Selection) {
			nodes := s.Find("a")
			link, _ := nodes.Attr("href")
			chap := nodes.Text()
			ExtractNumber := func(input string) string {
				// Define the regular expression pattern to match a number
				re := regexp.MustCompile(`\d+`)

				// Find the first numerical sequence in the input string
				match := re.FindString(input)

				return match
			}
			chap = ExtractNumber(chap)

			chapter = append(chapter, provider.Chapter{
				ChapterURl: link,
				Chapter:    chap,
			})

		})

		mangaDetail := &provider.MangaDetail{
			Title:       title,
			Status:      "ONGOING",
			ReleaseDate: release,
			Author:      "",
			Type:        typeHtml,
			// Sinopsis:    sinop,
			Chapter: chapter,
		}
		out <- mangaDetail
	})
}
