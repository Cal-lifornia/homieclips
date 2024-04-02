package frontend

import (
	"github.com/gin-gonic/gin"
	"homieclips/components"
	"net/http"
)

func CreateRootRoutes(rg *gin.RouterGroup) {
	rg.GET("", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", components.Layout(components.Home()))

	})
	rg.GET("/signin", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "", components.LoginPage())
	})

}
