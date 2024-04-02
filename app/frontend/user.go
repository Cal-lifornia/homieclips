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

var userPage = `
<div>
    <img class="avatar" src="{{ .picture }}"/>
    <h2>Welcome {{.nickname}}</h2>
</div>
`

func User(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	var result components.Profile
	err := mapstructure.Decode(profile, &result)
	if err != nil {
		log.Println("ran into error unmarshaling profile: ", err)
		ctx.JSON(http.StatusFailedDependency, util.ErrorResponse(err))
	}
	ctx.HTML(http.StatusOK, "", components.Layout(components.User(result)))
}
