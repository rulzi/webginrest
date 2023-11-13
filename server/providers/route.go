package providers

import (
	"context"
	"net/http"
	"webginrest/app/helpers"
	"webginrest/app/middleware"
	"webginrest/app/routers"
	"webginrest/server/config"

	"github.com/gin-gonic/gin"
)

// Route provide gin router
func Route(ctx context.Context) *gin.Engine {
	if !config.GetBoolean(`debug`) {
		gin.SetMode(gin.ReleaseMode)
	}

	// Creates a router without any middleware by default
	r := gin.New()

	// Notfound
	r.NoRoute(func(c *gin.Context) {
		message := "Page Not Found"
		result := helpers.ResponseData(http.StatusNotFound, nil, &message, nil, &message, nil)

		c.JSON(http.StatusNotFound, result)
	})

	// Middleware
	r = middleware.Middleware(ctx, r)

	// Routes Web
	r = routers.Web(ctx, r)
	r = routers.Guest(ctx, r)

	return r
}
