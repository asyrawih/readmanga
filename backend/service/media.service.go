package service

import (
	"context"

	"bacakomik/adapter"
	"bacakomik/record/entity"
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
	panic("not implemented") // TODO: Implement
}

// Retrive One Data
func (me *MediaService) GetOne(ctx context.Context, id int) *entity.Media {
	panic("not implemented") // TODO: Implement
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
	panic("not implemented") // TODO: Implement
}

// Update Data
func (me *MediaService) Update(ctx context.Context, data *entity.Media, id int) error {
	panic("not implemented") // TODO: Implement
}
