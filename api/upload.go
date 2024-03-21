package api

import (
	"fmt"
	db "homieclips/db/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server *Server) createUploadRoute(group *gin.RouterGroup) {
	recordings := group.Group("/upload")

	recordings.POST("", server.uploadRecording)
}

type uploadResponse struct {
	UploadComplete bool `json:"upload_complete,omitempty"`
}

func (server *Server) uploadRecording(ctx *gin.Context) {
	var response = uploadResponse{
		UploadComplete: false,
	}

	friendlyName := ctx.PostForm("friendly_name")
	gameName := ctx.PostForm("game_name")
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	objectName, err := uuid.NewV7()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	files := form.File["files"]

	for _, file := range files {
		server.queries.UploadRecording(ctx, objectName.String(), file, &response.UploadComplete)
		if ctx.IsAborted() {
			return
		}
	}
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

	ctx.JSON(http.StatusOK, response)
}
