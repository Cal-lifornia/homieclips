package frontend

import (
	"github.com/gin-gonic/gin"
	"homieclips/components"
	db "homieclips/db/models"
	"homieclips/util/gintemplrenderer"
	"net/http"
)

type Frontend struct {
	router *gin.Engine
	models *db.Models
}

func Init(router *gin.Engine, models *db.Models) {
	ginHtmlRenderer := router.HTMLRender

	router.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	frontend := &Frontend{
		router: router,
		models: models,
	}

	routerGroup := frontend.router.Group("")
	routerGroup.GET("/signin", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", components.LoginPage())
	})

	frontend.createAuthRoutes(routerGroup)
}
