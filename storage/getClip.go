package storage

import (
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

func (queries *Queries) StreamClip(ctx *gin.Context, objectName string) (*url.URL, error) {
	reqParams := make(url.Values)
	reqParams.Set("Content-Type", "video/mp4")
	reqParams.Set("Connection", "keep-alive")

	preSignedURL, err := queries.client.PresignedGetObject(
		ctx,
		queries.config.BucketName,
		objectName,
		time.Hour,
		reqParams,
	)
	if err != nil {
		return nil, err
	}
	return preSignedURL, nil
}
