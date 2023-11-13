package routers

import (
	"context"
	"webginrest/app/routers/resources/login"

	"github.com/gin-gonic/gin"
)

// Guest Guest router
func Guest(ctx context.Context, r *gin.Engine) *gin.Engine {

	loginController := login.Controller()

	g := r.Group("/")
	g.POST("/login", loginController.PostLogin)

	return r
}
