package api

import (
	db "homieclips/db/models"
	"homieclips/storage"
	"homieclips/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config  util.Config
	models  *db.Models
	queries *storage.Queries
	router  *gin.Engine
}

func NewServer(config util.Config, models *db.Models, queries *storage.Queries) *Server {
	server := &Server{
		config:  config,
		models:  models,
		queries: queries,
	}

	server.SetupRouter()

	return server
}

func (server *Server) SetupRouter() {
	router := gin.Default()
	public := router.Group("/api")
	public.GET("ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "pong"}) })

	server.createUploadRoute(public)
	server.createStreamRoute(public)
	server.createRecordingsRoutes()

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"message": err.Error(),
	}
}
