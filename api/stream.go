package api

import "github.com/gin-gonic/gin"

func (server *Server) createStreamRoute(group *gin.RouterGroup) {
	stream := group.Group("/stream")

	stream.GET("", server.streamRecording)
}

func (server *Server) streamRecording(ctx *gin.Context) {

}
