package service

import (
	"context"

	"bacakomik/record/entity"
	"bacakomik/repository/mysql"
)

type MangaService struct {
	repo *mysql.MangaRepository
}

// NewMangaService function
func NewMangaService(repo *mysql.MangaRepository) *MangaService {
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
	panic("not implemented") // TODO: Implement
}

// Retrive One Data
func (ma *MangaService) GetOne(ctx context.Context, id int) *entity.Manga {
	panic("not implemented") // TODO: Implement
}

// Update Data
func (ma *MangaService) Update(ctx context.Context, manga *entity.Manga, id int) error {
	panic("not implemented") // TODO: Implement
}
