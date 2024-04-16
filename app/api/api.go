package api

import (
	db "github.com/Cal-lifornia/homieclips/db/models"
	"github.com/Cal-lifornia/homieclips/storage"
	"github.com/Cal-lifornia/homieclips/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Api struct {
	router  *gin.Engine
	models  *db.Models
	storage *storage.Storage
	config  util.Config
}

func Init(router *gin.Engine, models *db.Models, storage *storage.Storage, config util.Config) {
	api := Api{
		router:  router,
		models:  models,
		storage: storage,
		config:  config,
	}

	routeGroup := router.Group("/api")
	routeGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	//routeGroup.Use(authenticator.IsAuthenticated())

	api.router.GET("/logout", api.logOut)

	api.createClipsRoute(routeGroup)
	api.createStreamRoute(routeGroup)
	api.createUploadRoute(routeGroup)
	routeGroup.Any("/storage/*proxyPath", api.proxy)
}
