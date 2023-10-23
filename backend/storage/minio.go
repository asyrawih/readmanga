package storage

import (
	"context"
	"errors"
	"net/http"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"
)

type MiniStorageServer struct {
	Host      string
	accessKey string
	secretKey string
}

// NewMinioStorageServer function
func NewMinioStorageServer(host string, AccessKey string, SecretKey string) *MiniStorageServer {
	return &MiniStorageServer{
		Host:      host,
		accessKey: AccessKey,
		secretKey: SecretKey,
	}
}

type Store struct {
	client     *minio.Client
	BucketName string
	Objectname string
}

// SetBucketName method
func (s *Store) SetBucketName(name string) {
	s.BucketName = name
}

// SetObjectName method
func (s *Store) SetObjectName(name string) {
	s.Objectname = name
}

// NewStore method
func (m *MiniStorageServer) NewStore() (*Store, error) {
	client, err := minio.New(m.Host, &minio.Options{
		Creds:  credentials.NewStaticV4(m.accessKey, m.secretKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Err(err).Msg("")
		return nil, err
	}
	return &Store{client: client, BucketName: "", Objectname: ""}, nil
}

// Upload method
func (m *Store) Upload(r *http.Response) error {
	_, err := m.client.PutObject(
		context.Background(),
		m.BucketName,
		m.Objectname,
		r.Body,
		r.ContentLength,
		minio.PutObjectOptions{},
	)
	if err != nil {
		log.Err(err).Msg("")
		return errors.New("error while uploading the image to MinIO")
	}

	return nil
}
