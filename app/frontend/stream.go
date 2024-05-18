package frontend

import (
	"fmt"
	"github.com/Cal-lifornia/homieclips/components"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (frontend *Frontend) getVideo(ctx *gin.Context) {

	objectName := ctx.Param("object_name")

	clipURL := fmt.Sprintf("https://%s/stream/%s/%s_master.m3u8", frontend.cloudfrontURL, objectName, objectName)

	ctx.HTML(http.StatusOK, "", components.Page(components.Video(clipURL), ctx.Copy()))
}
