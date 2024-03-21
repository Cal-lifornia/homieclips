package storage

import (
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func (queries *Queries) UploadRecording(ctx *gin.Context, objectName string, file *multipart.FileHeader, uploadComplete *bool) {
	contentType := file.Header["Content-Type"][0]
	openFile, err := file.Open()
	if err != nil {
		ctx.AbortWithError(http.StatusFailedDependency, err)
		return
	}

	uploadInfo, err := queries.client.PutObject(
		ctx,
		queries.config.BucketName,
		objectName,
		openFile,
		file.Size,
		minio.PutObjectOptions{ContentType: contentType},
	)

	if err != nil {
		ctx.AbortWithError(http.StatusFailedDependency, err)
		return
	}

	log.Print("file uploaded: ", uploadInfo)
	*uploadComplete = true
}
