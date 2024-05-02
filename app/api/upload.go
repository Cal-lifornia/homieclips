package api

import (
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
	UploadComplete bool `json:"upload_complete,omitempty"`
}

type uploadFileForm struct {
	FriendlyName string `json:"friendly_name" bson:"friendly_name" form:"friendly_name" binding:"required"`
	GameName     string `json:"game_name" bson:"game_name" form:"game_name" binding:"required"`
}

func (api *Api) uploadRecording(ctx *gin.Context) {
	var response = uploadResponse{
		UploadComplete: false,
	}

	var form uploadFileForm

	err := ctx.ShouldBind(&form)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
	}

	objectName, err := uuid.NewV7()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	recording := db.Clip{
		ObjectName:   objectName.String(),
		FriendlyName: form.FriendlyName,
		GameName:     form.GameName,
		Ready:        false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	upload := db.Upload{ObjectName: objectName.String(), Ready: false, CreatedAt: time.Now(), UpdatedAt: time.Now()}

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

	ctx.JSON(http.StatusOK, response)
}
