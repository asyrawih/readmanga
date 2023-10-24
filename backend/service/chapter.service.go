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
