package komikcast

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/iain17/go-cfscrape"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"

	"bacakomik/record/entity"
	"bacakomik/repository"
	"bacakomik/repository/mysql"
	"bacakomik/storage"
	"bacakomik/tools/provider"
)

func ProcessReadChapter(sizeWorker int) {
	defer func() {
		if r := recover(); r != nil {
			log.Err(errors.New("[ProcessReadChapter]: got panic try to restore it")).Msg("")
		}
	}()
	detail, err := LoadReadMangaDetil()
	if err != nil {
		log.Err(err).Msg("")
	}

	var DSN = "postgres://postgres:postgres@localhost:5432/readmanga"
	connect, err := repository.Connect(context.Background(), DSN)
	if err != nil {
		log.Err(err).Msg("error init database on crawler")
	}

	mr := mysql.NewMangaRepository(connect)
	cr := mysql.NewChapterRepository(connect)

	workerPool := make(chan struct{}, sizeWorker)
	defer close(workerPool)

	for _, md := range detail {
		title := strings.Trim(md.Title, "\n")
		sanitizeTitle := strings.TrimSpace(title)

		id, err := mr.Create(context.Background(), &entity.Manga{
			Title:        sanitizeTitle,
			Status:       strings.TrimSpace(md.Status),
			ReleaseDate:  md.ReleaseDate,
			TotalChapter: 0,
			Author:       "unknown",
			Type:         strings.TrimSpace(md.Type),
			Sinopsis:     strings.TrimSpace(md.Sinopsis),
			CreatedBy:    -1,
			CreatedAt:    time.Now(),
		})
		if err != nil {
			log.Err(err).Msg("")
		}
		for _, chapter := range md.Chapter {
			if chapter.ChapterURl != "" {
				cID, err := cr.Create(context.Background(), &entity.Chapter{
					MangaID: id,
					Chapter: chapter.Chapter,
					Content: strings.TrimSpace(md.Title),
				})
				if err != nil {
					log.Err(err).Msg("")
				}
				workerPool <- struct{}{}
				log.Info().Msgf("worker pool 20 from %v", len(workerPool))
				go GetChapterDetailImage(chapter.ChapterURl, connect, cID, workerPool)
			}
		}
	}
}

// GetChapterDetailImage function
func GetChapterDetailImage(chapterURL string, conn *pgxpool.Pool, cID int, worker <-chan struct{}) {
	defer func() {
		<-worker
		if r := recover(); r != nil {
			log.Err(errors.New("panice will recover")).Msg("")
		}
	}()
	media := mysql.NewMediaRepository(conn)
	r, err := cfscrape.Get(chapterURL)
	if err != nil {
		log.Err(err).Msg("")
	}

	defer r.Body.Close()

	doc, err := goquery.NewDocumentFromReader(r.Body)
	if err != nil {
		log.Err(err).Msg("")
	}

	doc.Find(".chapter_body").Each(func(_ int, s *goquery.Selection) {
		s.Find(".main-reading-area img").Each(func(_ int, s *goquery.Selection) {
			imagesURL, _ := s.Attr("src")
			komikcastURL := GetImage(imagesURL)
			media.Create(context.Background(), &entity.Media{
				ModelType: "chapters",
				ModelID:   cID,
				URL:       komikcastURL,
			})
		})
	})
}

func GetImage(imageURL string) string {
	req, err := http.NewRequest("GET", imageURL, nil)
	if err != nil {
		fmt.Println("Error while creating the request:", err)
		return ""
	}

	req.Header.Set("sec-ch-ua", `"Not=A?Brand";v="99", "Chromium";v="118"`)
	req.Header.Set("Referer", "https://komikcast.ch/")
	req.Header.Set("DNT", "1")
	req.Header.Set("sec-ch-ua-mobile", "?1")
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36",
	)
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while fetching the image:", err)
	}
	defer resp.Body.Close()

	// Check if the response was successful
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch the image:", resp.Status)
	}

	mss := storage.NewMinioStorageServer(
		"localhost:9000",
		"LR9boPTBwdBmeQzVHoCO",
		"7OVcEt2zky9sxs0GwtDEXsJdgVRPshBjEw6IwpBW",
	)
	s, err := mss.NewStore()
	if err != nil {
		fmt.Println("fail create new store please check your configuration")
	}

	f := strings.Split(imageURL, "/")
	filename := fmt.Sprintf("/sektekomik/%s/%s/%s", f[6], f[7], f[8])

	s.SetBucketName("manga")
	s.SetObjectName(filename)
	s.Upload(resp)

	log.Info().Msgf("Image downloaded and saved successfully! url:%s", filename)
	return filename

}

// LoadReadMangaDetil function
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
