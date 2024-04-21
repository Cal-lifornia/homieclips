package app

import (
	"encoding/gob"
	"log"

	"github.com/Cal-lifornia/homieclips/app/api"
	"github.com/Cal-lifornia/homieclips/app/authenticator"
	"github.com/Cal-lifornia/homieclips/app/frontend"
	db "github.com/Cal-lifornia/homieclips/db/models"
	"github.com/Cal-lifornia/homieclips/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config util.Config
	models *db.Models
	router *gin.Engine
	auth   *authenticator.Authenticator
}

func NewServer(config util.Config, models *db.Models) *Server {
	server := &Server{
		config: config,
		models: models,
	}

	server.SetupRouter()

	return server
}

func (server *Server) SetupRouter() {
	router := gin.Default()

	router.MaxMultipartMemory = 2000 << 20

	var err error

	server.auth, err = authenticator.New(server.config)
	if err != nil {
		log.Fatalf("failed to setup authenticator: %s\n", err)
	}

	err = router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatalf("failed to unset proxy: %s\n", err)
	}

	gob.Register(map[string]interface{}{})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("auth-session", store))
	baseUrl := router.Group("/")

	baseUrl.GET("/login", server.auth.Login)
	baseUrl.GET("/callback", server.auth.Callback)

	frontend.Init(router, server.models, server.config.CloudFrontURL)
	baseUrl.Static("/assets", "assets")

	api.Init(router, server.models, server.config)

	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
