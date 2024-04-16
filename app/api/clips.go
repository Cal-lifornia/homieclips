package api

import (
	"github.com/Cal-lifornia/homieclips/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Api) createClipsRoute(group *gin.RouterGroup) {
	clips := group.Group("/clips")

	clips.GET("", api.getClips)
	clips.GET(":object_name", api.getClip)
	clips.DELETE(":object_name", api.deleteClip)
}

func (api *Api) getClips(ctx *gin.Context) {
	clips, err := api.models.GetClips()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, util.Response{
		Results: clips,
		Count:   len(clips),
	})
}

func (api *Api) getClip(ctx *gin.Context) {
	objectName := ctx.Param("object_name")
	clip, err := api.models.GetClip(objectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, clip)
}

func (api *Api) deleteClip(ctx *gin.Context) {
	objectName := ctx.Param("object_name")
	err := api.models.DeleteClip(objectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "clip deleted successfully"})
}
