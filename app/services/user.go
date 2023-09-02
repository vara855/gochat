package services

import "github.com/vara855/gochat/app/models"

type UserService interface {
	Insert(user *models.User) (*models.User, error);

	Get(username *string) (*models.User, error);

	GetAll() ([]*models.User, error);
}