package controller

import (
	"net/http"
	"webginrest/app/helpers"

	"github.com/gin-gonic/gin"
)

// Index Index Controller
type Index interface {
	GetIndex(ctx *gin.Context)
}

type controller struct {
}

// IndexController init controller
func IndexController() Index {
	return &controller{}
}

func (con *controller) GetIndex(c *gin.Context) {
	//render with master
	var messageInt interface{}
	messageInt = "Hallo"
	result := helpers.ResponseData(http.StatusOK, &messageInt, nil, nil, nil, nil)

	c.JSON(http.StatusOK, result)
}
