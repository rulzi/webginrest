package controller

import (
	"net/http"
	"webginrest/app/helpers"
	formvalidation "webginrest/app/pkg/form_validation"
	uc "webginrest/app/pkg/usecase/login"

	"github.com/gin-gonic/gin"
)

// Login Login Controller
type Login interface {
	PostLogin(c *gin.Context)
}

type controller struct {
	usecase uc.Login
}

// LoginController init controller
func LoginController(usecase uc.Login) Login {
	return &controller{
		usecase: usecase,
	}
}

func (con *controller) PostLogin(c *gin.Context) {
	valErr := helpers.ValidationErrors(c, &formvalidation.Login{})

	if valErr != nil {
		var valInt interface{}
		valInt = valErr
		message := "Bad Request"
		result := helpers.ResponseData(http.StatusBadRequest, nil, &message, &valInt, &message, &valInt)
		c.JSON(http.StatusBadRequest, result)
		return
	}

	token, err := con.usecase.Login(c)

	if err != nil {
		var valInt interface{}
		valInt = err
		message := "Bad Request"
		result := helpers.ResponseData(http.StatusBadRequest, nil, &message, &valInt, &message, &valInt)
		c.JSON(http.StatusBadRequest, result)
		return
	}

	var tokenInt interface{}
	tokenInt = token

	result := helpers.ResponseData(http.StatusOK, &tokenInt, nil, nil, nil, nil)

	c.JSON(http.StatusOK, result)
}
