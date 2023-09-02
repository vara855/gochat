package services

import (
	"github.com/vara855/gochat/app/database"
	"github.com/vara855/gochat/app/models"
)

type config struct {
	memory_db database.MemoryDb
}

func (s *config) Insert(user *models.User) (*models.User, error) {
	s.memory_db.Insert(user)
	return user, nil
}

func (s *config) Get(username *string) (*models.User, error) {
	return s.memory_db.Get(username)
}

func (s *config) GetAll() ([]*models.User, error) {

	return s.memory_db.GetAll()
}

func New() UserService {
	return &config{
		memory_db: database.New(),
	}
}
