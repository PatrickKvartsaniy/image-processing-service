package storage

import (
	"bytes"
	gs "cloud.google.com/go/storage"
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"io"
)

type (
	Config struct {
		BucketName string
	}
	Storage struct {
		client *gs.Client
		bucket *gs.BucketHandle
	}
)

func New(ctx context.Context, cfg Config) (*Storage, error) {
	client, err := gs.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	return &Storage{
		client: client,
		bucket: client.Bucket(cfg.BucketName),
	}, nil
}

func (s Storage) Read(ctx context.Context, path string) (io.Reader, error) {
	obj := s.bucket.Object(path)
	r, err := obj.NewReader(ctx)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	if _, err := io.Copy(&b, r); err != nil {
		return nil, err
	}
	if err := r.Close(); err != nil {
		logrus.Error(err)
	}
	return &b, nil
}

func (s Storage) Upload(ctx context.Context, data io.Reader, extension string) (string, string, error) {
	obj := s.bucket.Object(uuid.NewV4().String() + extension)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, data); err != nil {
		return "", "", err
	}
	if err := w.Close(); err != nil {
		logrus.Error(err)
	}
	attr, err := obj.Attrs(ctx)
	if err != nil {
		return "", "", err
	}
	return attr.Name, attr.MediaLink, nil
}

func (s Storage) Close() {
	if err := s.client.Close(); err != nil {
		logrus.Error(err)
	}
}
