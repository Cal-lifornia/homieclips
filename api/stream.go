package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) createStreamRoute(group *gin.RouterGroup) {
	stream := group.Group("/stream")

	stream.GET(":object_name", server.getClipURL)
}

type getClipURLResponse struct {
	URL string `json:"url"`
}

func (server *Server) getClipURL(ctx *gin.Context) {
	objectName := ctx.Param("object_name")

	url, err := server.queries.StreamClip(ctx, objectName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, getClipURLResponse{
		URL: url.String(),
	})
}
