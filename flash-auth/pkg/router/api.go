package router

import (
	"net/http"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	// gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	// "repo.kmklabs.com/kmk-online/hermes/pkg/handler/endpoint/ad_tags"
	// "repo.kmklabs.com/kmk-online/hermes/pkg/handler/endpoint/deprecated_ads_tags"
	// "repo.kmklabs.com/kmk-online/hermes/pkg/handler/endpoint/whisper"
)

func setCORSHeaders(context *gin.Context) {
	if context.Request != nil && context.Request.Header != nil && context.Request.Header["Origin"] != nil {
		context.Writer.Header().Set("Access-Control-Allow-Origin", context.Request.Header["Origin"][0])
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}

func revisionHandler(c *gin.Context) {
	rev := os.Getenv("IMAGE_TAG")
	c.String(http.StatusOK, rev)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(func(context *gin.Context) {
		setCORSHeaders(context)
	})

	// router.Use(gintrace.Middleware("hermes-api"))
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// router.GET("/api/playlist/:videoId", whisper.Handle)
	// router.GET("/api/tags", ad_tags.Handle)
	// router.GET("/api/ads_params/video/:videoId", deprecated_ads_tags.Handle)
	router.GET("/revision", revisionHandler)
	router.GET("/healthcheck", func(context *gin.Context) {
		context.Status(http.StatusOK)
	})
	router.GET("/api/users/", user.List)
	router.GET("/api/users/verify/:session", user.Verify)
	router.POST("/api/users", user.CreateOne)

	return router
}
