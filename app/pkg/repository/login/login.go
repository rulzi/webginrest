package login

import (
	"webginrest/app/pkg/models"
	"webginrest/server/services"

	"github.com/jinzhu/gorm"
)

// Login Login Repository
type Login interface {
	GetUser(username string) (*models.User, error)
}

type repository struct {
	db *gorm.DB
}

// Repository init usecase
func Repository() Login {
	return &repository{
		services.DB,
	}
}

func (repo *repository) GetUser(username string) (*models.User, error) {
	var user models.User

	err := repo.db.Where("username = ?", username).
		Find(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
