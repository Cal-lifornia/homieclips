package frontend

import (
	"github.com/Cal-lifornia/homieclips/app/authenticator"
	"github.com/Cal-lifornia/homieclips/components"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (frontend *Frontend) createAuthRoutes(rg *gin.RouterGroup) {
	auth := rg.Group("", authenticator.IsAuthenticated())

	auth.GET("", frontend.homePage)
	auth.GET("/user", frontend.userPage)
	auth.GET("/stream/:object_name", frontend.getVideo)
	auth.GET("/upload-page", frontend.uploadPage)
}

func (frontend *Frontend) homePage(ctx *gin.Context) {

	clips, err := frontend.models.GetClips()
	if err != nil {
		ctx.HTML(http.StatusFailedDependency, "", components.Error(err))
		return
	}

	ctx.HTML(http.StatusOK, "", components.Page(components.Home(clips), ctx.Copy()))
}

func (frontend *Frontend) uploadPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "", components.Page(components.UploadPage(), ctx.Copy()))
}
