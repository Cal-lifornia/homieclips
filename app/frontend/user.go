package frontend

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"homieclips/components"
	"homieclips/util"
	"log"
	"net/http"
)

func (frontend *Frontend) userPage(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	var result components.Profile
	err := mapstructure.Decode(profile, &result)
	if err != nil {
		log.Println("ran into error unmarshaling profile: ", err)
		ctx.JSON(http.StatusFailedDependency, util.ErrorResponse(err))
	}
	ctx.HTML(http.StatusOK, "", components.Page(components.User(result)))
}
