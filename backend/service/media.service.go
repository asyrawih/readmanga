package service

import (
	"context"

	"bacakomik/adapter"
	"bacakomik/record/entity"

	"github.com/rs/zerolog/log"
)

type MediaService struct {
	repo adapter.RepoMediaCreational
}

// NewMediaService function
// Create New Instance for media
// For Interacting with media
func NewMediaService(repo adapter.RepoMediaCreational) *MediaService {
	return &MediaService{
		repo: repo,
	}
}

// Get All Data
func (me *MediaService) GetAll(ctx context.Context) []*entity.Media {
	m := me.repo.GetAll(ctx)
	return m
}

// Retrive One Data
func (me *MediaService) GetOne(ctx context.Context, id int) *entity.Media {
	m := me.repo.GetOne(ctx, id)
	return m
}

// Create Data
//
// create media with accept
func (me *MediaService) Create(ctx context.Context, data *entity.Media) (int, error) {
	i, err := me.repo.Create(ctx, data)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// Delete the record
func (me *MediaService) Delete(ctx context.Context, id int) bool {
	b := me.repo.Delete(ctx, id)
	return b
}

// Update Data
func (me *MediaService) Update(ctx context.Context, data *entity.Media, id int) error {
	if err := me.repo.Update(ctx, data, id); err != nil {
		log.Err(err).Msg("")
		return err
	}
	return nil
}
