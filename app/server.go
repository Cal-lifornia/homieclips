package app

import (
	"encoding/gob"
	"homieclips/app/api"
	"homieclips/app/authenticator"
	"homieclips/app/frontend"
	db "homieclips/db/models"
	"homieclips/storage"
	"homieclips/util"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config  util.Config
	models  *db.Models
	storage *storage.Storage
	router  *gin.Engine
	auth    *authenticator.Authenticator
}

func NewServer(config util.Config, models *db.Models, storage *storage.Storage) *Server {
	server := &Server{
		config:  config,
		models:  models,
		storage: storage,
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

	frontend.Init(router, server.models)
	baseUrl.Static("/assets", "assets")

	api.Init(router, server.models, server.storage, server.config)

	server.router = router

}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
