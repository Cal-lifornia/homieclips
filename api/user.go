package api

import (
	"encoding/json"
	"fmt"
	"homieclips/components"
	"log"
	"net/http"
	"net/url"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (server *Server) User(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	var result map[string]interface{}

	err := json.Unmarshal([]byte(fmt.Sprintf("%v", profile)), &result)
	if err != nil {
		log.Printf("ran into error unmarshaling profile: %s\n", err)
		ctx.JSON(http.StatusFailedDependency, errorResponse(err))
	}

	ctx.HTML(http.StatusOK, "", components.Layout(components.User(result)))
}

func (server *Server) LogOut(ctx *gin.Context) {
	logoutUrl, err := url.Parse("https://" + server.config.Auth0Domain + "/v2/logout")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + ctx.Request.Host)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", server.config.Auth0ClientID)
	logoutUrl.RawQuery = parameters.Encode()

	ctx.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}
