package middleware

import (
	"net/http"
	"strings"
	"webginrest/app/helpers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func globalRecover(c *gin.Context) {
	defer func(c *gin.Context) {
		if rec := recover(); rec != nil {

			message := "Error Server"
			result := helpers.ResponseData(http.StatusNotFound, nil, &message, &rec, &message, &rec)
			// that recovery also handle XHR's
			// you need handle it
			if XHR(c) {

				c.JSON(http.StatusInternalServerError, result)
			} else {
				logrus.Fatal(rec)
				c.JSON(http.StatusInternalServerError, result)
			}
		}
	}(c)
	c.Next()
}

// XHR xmlhttprequest
func XHR(c *gin.Context) bool {
	return strings.ToLower(c.Request.Header.Get("X-Requested-With")) == "xmlhttprequest"
}
