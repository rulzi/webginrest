package login

import (
	"errors"
	"webginrest/app/helpers"
	repo "webginrest/app/pkg/repository/login"
	"webginrest/app/utils/token"

	"github.com/gin-gonic/gin"
)

// Login Usecase
type Login interface {
	Login(c *gin.Context) (*string, error)
}

type usecase struct {
	repo repo.Login
}

// Usecase init usecase
func Usecase(repository repo.Login) Login {
	return &usecase{
		repo: repository,
	}
}

func (uc *usecase) Login(c *gin.Context) (*string, error) {
	username := c.PostForm("username")

	user, err := uc.repo.GetUser(username)

	if err != nil {
		errResponse := errors.New("User Not Found")
		return nil, errResponse
	}

	if err := helpers.ComparePassword(user.Password, c.PostForm("password")); err != nil {
		errResponse := errors.New("Password Not Match")
		return nil, errResponse
	}

	token, err := token.GenerateToken(user.ID)

	if err != nil {
		// errResponse := errors.New("Token Generate Failed")
		return nil, err
	}

	return &token, nil
}
