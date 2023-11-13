package routers

import (
	"context"
	"webginrest/app/middleware"
	"webginrest/app/routers/resources/index"

	"github.com/gin-gonic/gin"
)

// Web Web Router
func Web(ctx context.Context, r *gin.Engine) *gin.Engine {

	indexController := index.Controller()

	web := r.Group("/")
	web.Use(middleware.AuthMiddleware)
	web.GET("/", indexController.GetIndex)

	return r
}
