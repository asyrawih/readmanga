package service

import (
	"context"

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
func (ch *ChapterService) GetAll(ctx context.Context) []*entity.Chapter {
	c := ch.repo.GetAll(ctx)
	return c
}

// Retrive One Data
func (ch *ChapterService) GetOne(ctx context.Context, id int) *entity.Chapter {
	c := ch.repo.GetOne(ctx, id)
	return c
}

// Delete the record
func (ch *ChapterService) Delete(ctx context.Context, id int) bool {
	panic("not implemented") // TODO: Implement
}

// Update Data
func (ch *ChapterService) Update(ctx context.Context, data *entity.Chapter, id int) error {
	panic("not implemented") // TODO: Implement
}
