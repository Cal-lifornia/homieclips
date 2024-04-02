package api

import (
	"homieclips/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Api) createStreamRoute(group *gin.RouterGroup) {
	stream := group.Group("/stream")
	stream.GET(":object_name", api.getClipURL)
}

type getClipURLResponse struct {
	URL string `json:"url"`
}

func (api *Api) getClipURL(ctx *gin.Context) {
	objectName := ctx.Param("object_name")

	url, err := api.storage.StreamClip(ctx, objectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, getClipURLResponse{
		URL: url.String(),
	})
}
