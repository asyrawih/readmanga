package service

import (
	"context"

	"github.com/rs/zerolog/log"

	"bacakomik/adapter"
	"bacakomik/record/entity"
)

// ChapterService struct
type ChapterService struct {
	repo adapter.ChapterRepoCreational
}

// NewChapterService function
func NewChapterService(repo adapter.ChapterRepoCreational) *ChapterService {
	return &ChapterService{
		repo: repo,
	}
}

// Create Data
func (ch *ChapterService) Create(ctx context.Context, data *entity.Chapter) (int, error) {
	return ch.repo.Create(ctx, data)
}

// Get All Data
func (ch *ChapterService) GetAll(ctx context.Context) []*entity.ChapterWithMedia {
	var chm []*entity.ChapterWithMedia
	chap := ch.repo.GetAll(ctx)
	for _, c := range chap {
		cwm := new(entity.ChapterWithMedia)
		cwm.Chapter = *c
		chm = append(chm, cwm)
	}
	return chm
}

// Retrive One Data
func (ch *ChapterService) GetOne(ctx context.Context, id int) *entity.ChapterWithMedia {
	c := ch.repo.GetOne(ctx, id)
	cwm := new(entity.ChapterWithMedia)
	cwm.Chapter = *c
	medias := ch.repo.NewApi().GetMedias(ctx, id)
	log.Info().Msgf("%v", medias)
	for _, m := range medias {
		cwm.Medias = append(cwm.Medias, *m)
	}
	return cwm
}

// Delete the record
func (ch *ChapterService) Delete(ctx context.Context, id int) bool {
	panic("not implemented") // TODO: Implement
}

// Update Data
func (ch *ChapterService) Update(ctx context.Context, data *entity.Chapter, id int) error {
	panic("not implemented") // TODO: Implement
}
