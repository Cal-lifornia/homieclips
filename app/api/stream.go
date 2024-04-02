package api

import (
	"homieclips/app"
	"homieclips/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *app.Server) createStreamRoute(group *gin.RouterGroup) {
	stream := group.Group("/stream")
	stream.GET(":object_name", server.getClipURL)
}

type getClipURLResponse struct {
	URL string `json:"url"`
}

func (server *app.Server) getClipURL(ctx *gin.Context) {
	objectName := ctx.Param("object_name")

	url, err := server.storage.StreamClip(ctx, objectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, getClipURLResponse{
		URL: url.String(),
	})
}
