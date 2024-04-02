package api

import (
	"encoding/gob"
	"homieclips/api/authenticator"
	"homieclips/components"
	db "homieclips/db/models"
	"homieclips/storage"
	"homieclips/util"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config  util.Config
	models  *db.Models
	queries *storage.Queries
	router  *gin.Engine
	auth    *authenticator.Authenticator
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

	var err error

	server.auth, err = authenticator.New(server.config)
	if err != nil {
		log.Fatalf("failed to setup authenticator: %s\n", err)
	}

	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", components.LoginPage())
	})

	router.Static("/assets", "assets")

	router.GET("/login", server.auth.Login)
	router.GET("/callback", server.auth.Callback)
	router.GET("/user", authenticator.IsAuthenticated(), server.User)

	api := router.Group("/api")

	api.Use(authenticator.IsAuthenticated())

	api.Any("/storage/*proxyPath", server.proxy)

	api.GET("ping", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "pong"}) })

	server.createUploadRoute(api)
	server.createStreamRoute(api)
	server.createClipsRoute(api)

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
