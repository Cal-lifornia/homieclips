package app

import (
	"encoding/gob"
	"homieclips/app/authenticator"
	"homieclips/app/frontend"
	db "homieclips/db/models"
	"homieclips/storage"
	"homieclips/util"
	"homieclips/util/gintemplrenderer"
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

	ginHtmlRenderer := router.HTMLRender

	router.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}
	err = router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("failed to unset proxy: %s\n", err)
	}

	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	baseUrl := router.Group("/")
	baseUrl.Use(sessions.Sessions("auth-session", store))

	baseUrl.GET("/login", server.auth.Login)
	baseUrl.GET("/callback", server.auth.Callback)
	baseUrl.GET("/user", authenticator.IsAuthenticated(), frontend.User)
	baseUrl.GET("/logout", server.LogOut)

	frontend.CreateRootRoutes(baseUrl)
	baseUrl.Static("/assets", "assets")

	api := baseUrl.Group("/api")

	//api.Use(authenticator.IsAuthenticated())

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
