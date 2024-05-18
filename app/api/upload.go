package api

import (
	"mime/multipart"
	"net/http"
	"time"

	db "github.com/Cal-lifornia/homieclips/db/models"
	"github.com/Cal-lifornia/homieclips/util"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func (api *Api) createUploadRoute(group *gin.RouterGroup) {
	upload := group.Group("/upload")

	upload.POST("", api.uploadRecording)
}

type uploadResponse struct {
	PresignPutUrl string `json:"presign_put_url"`
}

type uploadFileForm struct {
	ObjectName   string                `json:"object_name" bson:"object_name" form:"object_name" binding:"required"`
	FriendlyName string                `json:"friendly_name" bson:"friendly_name" form:"friendly_name" binding:"required"`
	GameName     string                `json:"game_name" bson:"game_name" form:"game_name" binding:"required"`
	File         *multipart.FileHeader `form:"uploaded_file" binding:"required"`
}

func (api *Api) uploadRecording(ctx *gin.Context) {
	var form uploadFileForm

	err := ctx.ShouldBind(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	file, err := form.File.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	defer file.Close()

	buffer := make([]byte, 512)

	_, err = file.Read(buffer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	contentType := http.DetectContentType(buffer)
	if contentType != "video/mp4" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect video type, must be an MP4"})
		return
	}

	presignPutUrl, err := api.putPresignURL(form.ObjectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	recording := db.Clip{
		ObjectName:   form.ObjectName,
		FriendlyName: form.FriendlyName,
		GameName:     form.GameName,
		Ready:        false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	upload := db.Upload{ObjectName: form.ObjectName, Ready: false, CreatedAt: time.Now(), UpdatedAt: time.Now()}

	_, err = api.models.CreateClip(recording)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	_, err = api.models.CreateUpload(upload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, uploadResponse{PresignPutUrl: presignPutUrl})
}

type objectNameResponse struct {
	ObjectName string `json:"object_name"`
}

func (api *Api) generateObjectName(ctx *gin.Context) {
	objectName, err := uuid.NewV7()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, objectNameResponse{ObjectName: objectName.String()})
}
