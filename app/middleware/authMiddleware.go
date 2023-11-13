package middleware

import (
	"net/http"
	"webginrest/app/helpers"
	"webginrest/app/utils/token"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware Auth middleware
func AuthMiddleware(c *gin.Context) {
	err := token.TokenValid(c)
	if err != nil {
		message := "Unauthorized"
		result := helpers.ResponseData(http.StatusUnauthorized, nil, &message, nil, &message, nil)

		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
		return
	}
	c.Next()
}
