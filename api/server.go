package api

import (
	db "homieclips/db/models"
	"homieclips/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

type Server struct {
	config      util.Config
	models      *db.Models
	minioClient *minio.Client
	router      *gin.Engine
}

func NewServer(config util.Config, models *db.Models, minioClient *minio.Client) *Server {
	return &Server{
		config:      config,
		models:      models,
		minioClient: minioClient,
	}
}

func (server *Server) SetupRouter() {
	router := gin.Default()
	public := router.Group("/api")
	public.GET("ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "pong"}) })

	server.createRecordingsRoutes(public)

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
