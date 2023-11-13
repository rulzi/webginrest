package helpers

import (
	"errors"
	"log"
	"os"
	"webginrest/app/pkg/transformers"
	"webginrest/server/config"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// BaseURL get BaseUrl from config
func BaseURL() string {
	return config.Get(`appurl`)
}

// HashPassword Hash Password
func HashPassword(password string) string {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	return string(result)
}

// ComparePassword Comparing Password
func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}

// DumpAndDie Dump and Die
func DumpAndDie(data ...interface{}) {
	spew.Dump(data)

	os.Exit(1)
}

// FlashMessage flash massage gin
func FlashMessage(c *gin.Context, messages []string, vars ...string) {
	session := sessions.Default(c)
	for _, message := range messages {
		session.AddFlash(message, vars...)
	}
	if err := session.Save(); err != nil {
		log.Printf("error in flashMessage saving session: %s", err)
	}
}

// Flashes flash gin
func Flashes(c *gin.Context, vars ...string) []interface{} {
	session := sessions.Default(c)
	flashes := session.Flashes(vars...)
	if len(flashes) != 0 {
		if err := session.Save(); err != nil {
			log.Printf("error in flashes saving session: %s", err)
		}
	}
	return flashes
}

// GetErrorMsg Get error massage validator
func GetErrorMsg(field string, fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " should be email format"
	case "lte":
		return field + " should be less than " + fe.Param()
	case "gte":
		return field + " should be greater than " + fe.Param()
	case "eqfield":
		return field + " should be same as " + fe.Param()
	}
	return "Unknown error"
}

// ValidationErrors Cek Error Validator
func ValidationErrors(c *gin.Context, m interface{}) []string {
	err := c.ShouldBind(m)

	if err != nil {
		var ve validator.ValidationErrors
		var errMsg []string
		if errors.As(err, &ve) {

			for _, fe := range ve {
				errMsg = append(errMsg, GetErrorMsg(fe.Field(), fe))
			}

			return errMsg
		}
	}

	return nil
}

func ResponseData(status int, data *interface{}, messageTitle *string, userMessage *interface{}, errorMessage *string, err *interface{}) transformers.Result {

	message := transformers.UserMassage{
		Title:   messageTitle,
		Message: userMessage,
	}

	errorM := transformers.Error{
		Message: errorMessage,
		Reason:  err,
	}

	result := transformers.Result{
		Status:      status,
		Data:        data,
		UserMassage: &message,
		Error:       &errorM,
	}

	return result
}
