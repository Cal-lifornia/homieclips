package storage

import (
	"bytes"
	"fmt"
	"homieclips/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func (queries *Queries) GetFile(ctx *gin.Context, objectName string) {
	object, err := queries.client.GetObject(
		ctx,
		queries.config.BucketName,
		objectName,
		minio.GetObjectOptions{},
	)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	defer object.Close()

	fileInfo, err := object.Stat()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	buffer := make([]byte, fileInfo.Size)
	object.Read(buffer)

	ctx.Writer.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size))
	ctx.Writer.Header().Set("Content-Type", "video/mp4")
	ctx.Writer.Header().Set("Connection", "keep-alive")
	ctx.Writer.Header().Set("Content-Range", fmt.Sprintf("bytes 0-%d/%d", fileInfo.Size, fileInfo.Size))

	ctx.DataFromReader(200, fileInfo.Size, "video/mp4", bytes.NewReader(buffer), nil)
}
