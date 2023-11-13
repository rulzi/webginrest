package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
)

// Middleware default middleware
func Middleware(ctx context.Context, r *gin.Engine) *gin.Engine {

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.Use(globalRecover)

	return r
}
