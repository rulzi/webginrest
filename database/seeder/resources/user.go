package resources

import (
	"webginrest/app/helpers"
	"webginrest/app/pkg/models"
)

// User seeding user
func User() []*models.User {
	var users []*models.User

	users = append(users, &models.User{
		Email:    "admin@gmail.com",
		Name:     "admin",
		Username: "admin",
		Address:  "surabaya",
		Password: helpers.HashPassword("default12"),
	})

	return users
}
