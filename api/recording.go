package api

import (
	"fmt"
	db "homieclips/db/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server *Server) createRecordingsRoutes(group *gin.RouterGroup) {
	recordings := group.Group("/recordings")

	recordings.POST("", server.UploadRecording)
}

func (server *Server) UploadRecording(ctx *gin.Context) {
	friendlyName := ctx.PostForm("friendly_name")
	gameName := ctx.PostForm("game_name")
	/*
		form, err := ctx.MultipartForm()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	*/
	objectName, err := uuid.NewV7()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	/*
		files := form.File["files"]

		for _, file := range files {
			fileOpen, err := file.Open()
			contentType := file.Header["Content-Type"][0]
			if err != nil {
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}

			uploadInfo, err := server.minioClient.PutObject(
				context.TODO(),
				server.config.BucketName,
				objectName.String(),
				fileOpen,
				file.Size,
				minio.PutObjectOptions{ContentType: contentType},
			)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, errorResponse(err))
				return
			}
			fmt.Println("file uploaded successfully: ", uploadInfo)
		}
	*/
	recording := db.Recording{
		ObjectName:   objectName.String(),
		FriendlyName: friendlyName,
		GameName:     gameName,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	result, err := server.models.CreateRecording(recording)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	fmt.Println(result.InsertedID)

	ctx.JSON(http.StatusOK, gin.H{"message": "created record with id"})
}
