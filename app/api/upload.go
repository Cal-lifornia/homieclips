package api

import (
	"fmt"
	db "homieclips/db/models"
	"homieclips/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (api *Api) createUploadRoute(group *gin.RouterGroup) {
	upload := group.Group("/upload")

	upload.POST("", api.uploadRecording)
}

type uploadResponse struct {
	UploadComplete bool `json:"upload_complete,omitempty"`
}

func (api *Api) uploadRecording(ctx *gin.Context) {
	var response = uploadResponse{
		UploadComplete: false,
	}

	friendlyName := ctx.PostForm("friendly_name")
	gameName := ctx.PostForm("game_name")
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	objectName, err := uuid.NewV7()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	files := form.File["files"]

	for _, file := range files {
		api.storage.UploadClip(ctx, objectName.String(), file, &response.UploadComplete)
		if ctx.IsAborted() {
			return
		}
	}
	recording := db.Clip{
		ObjectName:   objectName.String(),
		FriendlyName: friendlyName,
		GameName:     gameName,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	result, err := api.models.CreateClip(recording)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	fmt.Println(result.InsertedID)

	ctx.JSON(http.StatusOK, response)
}
