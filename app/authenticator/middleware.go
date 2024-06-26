package authenticator

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// IsAuthenticated is a middleware that checks if
// the user has already been authenticated previously.
func IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if sessions.Default(ctx).Get("profile") == nil {
			ctx.Redirect(http.StatusSeeOther, "/signin")
		} else {
			ctx.Next()
		}
	}
}
