package storage

import (
	"homieclips/util"

	"github.com/minio/minio-go/v7"
)

type Queries struct {
	client *minio.Client
	config util.Config
}

func New(client *minio.Client, config util.Config) *Queries {
	return &Queries{
		client: client,
		config: config,
	}
}
