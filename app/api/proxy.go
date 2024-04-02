package api

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func (api *Api) proxy(ctx *gin.Context) {
	remoteUrl, err := url.Parse("https://" + api.config.MinioURL)
	if err != nil {
		log.Fatalf("failed to connect reverse proxy: %s", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remoteUrl)
	proxy.Director = func(req *http.Request) {
		req.Header = ctx.Request.Header
		req.Host = remoteUrl.Host
		req.URL.Scheme = "https"
		req.URL.Host = remoteUrl.Host
		req.URL.Path = ctx.Param("proxyPath")
	}

	proxy.ServeHTTP(ctx.Writer, ctx.Request)
}
