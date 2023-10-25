package service

import (
	"context"

	"github.com/rs/zerolog/log"

	"bacakomik/adapter"
	"bacakomik/record/entity"
	"bacakomik/record/model"
)

type MangaService struct {
	repo adapter.RepoMangaCreational
}

// NewMangaService function
func NewMangaService(repo adapter.RepoMangaCreational) *MangaService {
	return &MangaService{
		repo: repo,
	}
}

// Create Data
func (ma *MangaService) Create(ctx context.Context, manga *entity.Manga) (int, error) {
	return ma.repo.Create(ctx, manga)
}

// Get All Data
func (ma *MangaService) GetAll(ctx context.Context) []*entity.Manga {
	m := ma.repo.GetAll(ctx)
	return m
}

// Retrive One Data
func (ma *MangaService) GetOne(ctx context.Context, id int) *entity.Manga {
	m := ma.repo.GetOne(ctx, id)
	return m
}

// GetMangaWithChapters method
func (ma *MangaService) GetMangaWithChapters(ctx context.Context, id int) model.GetMangaChapter {
	var mangaChapter model.GetMangaChapter
	m := ma.repo.GetOne(ctx, id)
	mangaChapter.Manga = *m
	c, err := ma.repo.NewApi().GetChapters(ctx, id)
	if err != nil {
		log.Err(err).Msg("")
	}
	mangaChapter.Chapters = c
	return mangaChapter
}

// Get Access to instance of of T
func (ma *MangaService) NewApi() *MangaService {
	return ma
}

// Update Data
func (ma *MangaService) Update(ctx context.Context, manga *entity.Manga, id int) error {
	if err := ma.repo.Update(ctx, manga, id); err != nil {
		return err
	}
	return nil
}

// Delete method
func (ma *MangaService) Delete(ctx context.Context, id int) bool {
	return ma.repo.Delete(ctx, id)
}
