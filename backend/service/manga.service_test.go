package service

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"bacakomik/adapter"
	mocks "bacakomik/mocks/adapter"
	"bacakomik/record/entity"

	"github.com/stretchr/testify/mock"
)

func setupCreateTest(t *testing.T, rmc *mocks.RepoMangaCreational) (*MangaService, *entity.Manga, context.Context) {
	sampleData := &entity.Manga{
		ID:           1,
		Title:        "Naruto",
		Status:       "ongoing",
		ReleaseDate:  "2023",
		TotalChapter: 200,
		Author:       "hanan",
		Type:         "manga",
		Sinopsis:     "tidak ada",
		CreatedBy:    2,
		CreatedAt:    time.Now(),
	}

	ctx := context.Background()

	return &MangaService{repo: rmc}, sampleData, ctx
}

func TestNewMangaService(t *testing.T) {
	rmc := mocks.NewRepoMangaCreational(t)
	type args struct {
		repo adapter.RepoMangaCreational
	}
	tests := []struct {
		name string
		args args
		want *MangaService
	}{
		{
			name: "Test Should Oke Inserting Data",
			args: args{
				repo: rmc,
			},
			want: &MangaService{
				repo: rmc,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMangaService(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMangaService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMangaService_Create(t *testing.T) {
	rmc := mocks.NewRepoMangaCreational(t)
	tests := []struct {
		name      string
		setup     func(rmc *mocks.RepoMangaCreational) (*MangaService, *entity.Manga, context.Context)
		wantError bool
	}{
		{
			name: "should success create manga",
			setup: func(rmc *mocks.RepoMangaCreational) (*MangaService, *entity.Manga, context.Context) {
				ms, m, ctx := setupCreateTest(t, rmc)
				rmc.On("Create", ctx, m).Return(102, nil)
				return ms, m, ctx
			},
			wantError: false,
		},
		{
			name: "should fail create manga if not giving correct args",
			setup: func(rmc *mocks.RepoMangaCreational) (*MangaService, *entity.Manga, context.Context) {
				ms, m, ctx := setupCreateTest(t, rmc)
				err := errors.New("error")
				rmc.On("Create", ctx, mock.Anything).Return(1, err)
				return ms, m, ctx
			},
			wantError: true,
		},
		{
			name: "by giving manga into args but return nil",
			setup: func(rmc *mocks.RepoMangaCreational) (*MangaService, *entity.Manga, context.Context) {
				ms, _, ctx := setupCreateTest(t, rmc)
				err := errors.New("error")
				rmc.On("Create", ctx, mock.Anything).Return(0, err)
				return ms, nil, ctx
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			ms, m, ctx := tt.setup(rmc)
			_, err := ms.Create(ctx, m)
			if (err != nil) != tt.wantError {
				t.Errorf("MangaService.Create() error = %v, wantErr %v", err, tt.wantError)
			}
		})
	}
}

func TestMangaService_GetAll(t *testing.T) {
	rmc := mocks.NewRepoMangaCreational(t)
	tests := []struct {
		name      string
		setup     func(rmc *mocks.RepoMangaCreational) (*MangaService, context.Context)
		wantError bool
	}{
		{
			name: "should return list of manga",
			setup: func(rmc *mocks.RepoMangaCreational) (*MangaService, context.Context) {
				ms, _, ctx := setupCreateTest(t, rmc)
				rmc.On("GetAll", ctx).Return([]*entity.Manga{})
				return ms, ctx
			},
			wantError: false,
		},
		{
			name: "should error if return a nil",
			setup: func(rmc *mocks.RepoMangaCreational) (*MangaService, context.Context) {
				ms, _, ctx := setupCreateTest(t, rmc)
				rmc.On("GetAll", ctx).Return(nil)
				return ms, ctx
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms, ctx := tt.setup(rmc)
			m := ms.GetAll(ctx)
			t.Log(m)
		})
	}

}

func TestMangaService_GetOne(t *testing.T) {
	var rmc *mocks.RepoMangaCreational
	tests := []struct {
		name      string
		setup     func(rmc *mocks.RepoMangaCreational) (*MangaService, context.Context)
		wantError bool
	}{
		{
			name: "should return list of manga",
			setup: func(rmc *mocks.RepoMangaCreational) (*MangaService, context.Context) {
				ms, _, ctx := setupCreateTest(t, rmc)
				rmc.On("GetOne", ctx, mock.Anything).Return(&entity.Manga{Title: "test"})
				return ms, ctx
			},
			wantError: false,
		},
		{
			name: "Return nil if not query set has found",
			setup: func(rmc *mocks.RepoMangaCreational) (*MangaService, context.Context) {
				ms, _, ctx := setupCreateTest(t, rmc)
				rmc.On("GetOne", ctx, 1).Return(nil)
				return ms, ctx
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rmc = mocks.NewRepoMangaCreational(t)
			ms, ctx := tt.setup(rmc)
			m := ms.GetOne(ctx, 1)
			if (m == nil) != tt.wantError {
				t.Error(m)
			}
		})
	}

}

func TestMangaService_Update(t *testing.T) {
	tests := []struct {
		name      string
		args      int
		setup     func(rmc *mocks.RepoMangaCreational) (*MangaService, *entity.Manga, context.Context)
		wantError bool
	}{
		{
			name: "Should not returning error",
			args: 1,
			setup: func(rmc *mocks.RepoMangaCreational) (*MangaService, *entity.Manga, context.Context) {
				ms, m, ctx := setupCreateTest(t, rmc)
				rmc.On("Update", ctx, m, mock.Anything).Return(nil)
				return ms, m, ctx
			},

			wantError: false,
		},
		{
			name: "should not returning error",
			args: 1,
			setup: func(rmc *mocks.RepoMangaCreational) (*MangaService, *entity.Manga, context.Context) {
				ms, m, ctx := setupCreateTest(t, rmc)
				rmc.On("Update", ctx, m, mock.Anything).Return(errors.New("error fail update data"))
				return ms, m, ctx
			},

			wantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(t.Name(), func(t *testing.T) {
			rmc := mocks.NewRepoMangaCreational(t)
			ms, m, ctx := tt.setup(rmc)
			err := ms.Update(ctx, m, tt.args)
			if (err != nil) != tt.wantError {
				t.Fatal(err)
			}
		})
	}
}

func TestMangaService_Delete(t *testing.T) {

	tests := []struct {
		name      string
		args      int
		setup     func(rmc *mocks.RepoMangaCreational) (*MangaService, context.Context)
		wantError bool
	}{
		{
			name: "should return true",
			args: 1,
			setup: func(rmc *mocks.RepoMangaCreational) (*MangaService, context.Context) {
				ms, _, ctx := setupCreateTest(t, rmc)
				rmc.On("Delete", ctx, mock.Anything).Return(false)
				return ms, ctx
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rmc := mocks.NewRepoMangaCreational(t)
			ms, ctx := tt.setup(rmc)
			b := ms.Delete(ctx, tt.args)
			if b != tt.wantError {
				t.Errorf("Expeted %v : actual = %v", tt.wantError, b)
			}
		})
	}

}
