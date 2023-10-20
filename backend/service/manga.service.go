package service

import (
	"context"

	"bacakomik/adapter"
	"bacakomik/record/entity"
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
func (ma *MangaService) Create(ctx context.Context, manga *entity.Manga) error {
	if err := ma.repo.Create(ctx, manga); err != nil {
		return err
	}
	return nil
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

// Update Data
func (ma *MangaService) Update(ctx context.Context, manga *entity.Manga, id int) error {
	if err := ma.repo.Update(ctx, manga, id); err != nil {
		return err
	}
	return nil
}

// Delete method
func (ma *MangaService) Delete(ctx context.Context, id int) bool {
	panic("not implemented") // TODO: Implement
}
