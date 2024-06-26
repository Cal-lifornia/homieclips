package frontend

import (
	"github.com/Cal-lifornia/homieclips/components"
	db "github.com/Cal-lifornia/homieclips/db/models"
	"github.com/Cal-lifornia/homieclips/util/gintemplrenderer"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Frontend struct {
	router        *gin.Engine
	models        *db.Models
	cloudfrontURL string
}

func Init(router *gin.Engine, models *db.Models, cloudFrontURL string) {
	ginHtmlRenderer := router.HTMLRender

	router.HTMLRender = &gintemplrenderer.HTMLTemplRenderer{FallbackHtmlRenderer: ginHtmlRenderer}

	frontend := &Frontend{
		router:        router,
		models:        models,
		cloudfrontURL: cloudFrontURL,
	}

	routerGroup := frontend.router.Group("")
	routerGroup.GET("/signin", func(ctx *gin.Context) {
		signedOut := func() bool {
			if sessions.Default(ctx).Get("profile") == nil {
				return false
			}
			return true
		}

		ctx.HTML(http.StatusOK, "", components.PageWithoutNav(components.ProfilePage(signedOut())))
	})

	frontend.createAuthRoutes(routerGroup)
}
