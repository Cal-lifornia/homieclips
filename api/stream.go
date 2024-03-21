package api

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) createStreamRoute(group *gin.RouterGroup) {
	stream := group.Group("/stream")

	stream.GET(":object_name", server.streamRecording)
}

func (server *Server) streamRecording(ctx *gin.Context) {
	objectName := ctx.Param("object_name")

	server.queries.GetFile(ctx, objectName)
}
