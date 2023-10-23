package komikcast

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/iain17/go-cfscrape"
	"github.com/rs/zerolog/log"

	"bacakomik/tools/provider"
)

func ProcessReadChapter() {
	detail, err := LoadReadMangaDetil()
	if err != nil {
		log.Err(err).Msg("")
	}

	for _, md := range detail {
		fmt.Printf("md: %v\n", md.Title)
		for _, chapter := range md.Chapter {
			if chapter.ChapterURl != "" {
				GetChapterDetailImage(chapter.ChapterURl)
			}
		}
	}

}

// GetChapterDetailImage function
func GetChapterDetailImage(chapterURL string) {
	r, err := cfscrape.Get(chapterURL)
	if err != nil {
		log.Err(err).Msg("")
	}

	defer r.Body.Close()

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Err(err).Msg("")
	}

	doc.Find(".main-reading-area").Each(func(_ int, s *goquery.Selection) {
		chapter_html, _ := s.Find("img").Attr("src")
		fmt.Printf("chapter_html: %v\n", chapter_html)
		GetImage(chapter_html)

	})
}

func GetImage(imageURL string) {
	outputFile := "./dataset/downloaded_image.jpg"

	req, err := http.NewRequest("GET", imageURL, nil)
	if err != nil {
		fmt.Println("Error while creating the request:", err)
		return
	}

	req.Header.Set("sec-ch-ua", `"Not=A?Brand";v="99", "Chromium";v="118"`)
	req.Header.Set("Referer", "https://komikcast.ch/")
	req.Header.Set("DNT", "1")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
	)
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while fetching the image:", err)
		return
	}
	defer resp.Body.Close()

	// Check if the response was successful
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch the image:", resp.Status)
		return
	}

	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error while creating the file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error while saving the image:", err)
		return
	}

	fmt.Println("Image downloaded and saved successfully!")
}

func LoadReadMangaDetil() ([]*provider.MangaDetail, error) {
	b, err := os.ReadFile("./dataset/manga-detail.json")
	if err != nil {
		return nil, err
	}

	var mangaDetail []*provider.MangaDetail

	r := strings.NewReader(string(b))
	d := json.NewDecoder(r)

	if err := d.Decode(&mangaDetail); err != nil {
		return nil, errors.New("fail decode manga detail ")
	}

	return mangaDetail, nil
}
