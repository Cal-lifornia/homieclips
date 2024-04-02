package storage

import (
	"homieclips/util"

	"github.com/minio/minio-go/v7"
)

type Storage struct {
	client *minio.Client
	config util.Config
}

func New(client *minio.Client, config util.Config) *Storage {
	return &Storage{
		client: client,
		config: config,
	}
}
