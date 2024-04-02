package frontend

import (
	"github.com/gin-gonic/gin"
	"homieclips/app/authenticator"
	"homieclips/components"
	"net/http"
)

func (frontend *Frontend) createAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("")

	auth.GET("", authenticator.IsAuthenticated(), frontend.homePage)
	auth.GET("/user", authenticator.IsAuthenticated(), frontend.userPage)

}

func (frontend *Frontend) homePage(ctx *gin.Context) {
	clips, err := frontend.models.GetClips()
	if err != nil {
		ctx.HTML(http.StatusFailedDependency, "", components.Page(components.Error(err)))
		return
	}

	ctx.HTML(http.StatusOK, "", components.Page(components.Home(clips)))
}
