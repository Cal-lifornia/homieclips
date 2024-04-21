package frontend

import (
	"github.com/Cal-lifornia/homieclips/app/authenticator"
	"github.com/Cal-lifornia/homieclips/components"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (frontend *Frontend) createAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("")

	auth.GET("", authenticator.IsAuthenticated(), frontend.homePage)
	auth.GET("/user", authenticator.IsAuthenticated(), frontend.userPage)
	auth.GET("/stream/:object_name", authenticator.IsAuthenticated(), frontend.getVideo)
}

func (frontend *Frontend) homePage(ctx *gin.Context) {
	clips, err := frontend.models.GetClips()
	if err != nil {
		ctx.HTML(http.StatusFailedDependency, "", components.Page(components.Error(err)))
		return
	}

	ctx.HTML(http.StatusOK, "", components.Page(components.Home(clips)))
}
