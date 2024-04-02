package api

import (
	"homieclips/app"
	"homieclips/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *app.Server) createClipsRoute(group *gin.RouterGroup) {
	clips := group.Group("/clips")

	clips.GET("", server.getClips)
	clips.GET(":object_name", server.getClip)
	clips.DELETE(":object_name", server.deleteClip)
}

func (server *app.Server) getClips(ctx *gin.Context) {
	clips, err := server.models.GetClips()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, util.Response{
		Results: clips,
		Count:   len(clips),
	})
}

func (server *app.Server) getClip(ctx *gin.Context) {
	objectName := ctx.Param("object_name")
	clip, err := server.models.GetClip(objectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, clip)
}

func (server *app.Server) deleteClip(ctx *gin.Context) {
	objectName := ctx.Param("object_name")
	err := server.models.DeleteClip(objectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "clip deleted successfully"})
}
