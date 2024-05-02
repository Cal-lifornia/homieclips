package api

import (
	"net/http"

	db "github.com/Cal-lifornia/homieclips/db/models"
	"github.com/Cal-lifornia/homieclips/util"

	"github.com/gin-gonic/gin"
)

type Api struct {
	router *gin.Engine
	models *db.Models
	config util.Config
}

func Init(router *gin.Engine, models *db.Models, config util.Config) {
	api := Api{
		router: router,
		models: models,
		config: config,
	}

	routeGroup := router.Group("/api")
	routeGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	//routeGroup.Use(authenticator.IsAuthenticated())

	api.router.GET("/logout", api.logOut)

	api.createClipsRoute(routeGroup)
	api.createUploadRoute(routeGroup)
}
