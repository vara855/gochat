package database

import (
	"errors"
	"fmt"
	"sync"

	"github.com/vara855/gochat/app/models"
)

type userStore []*models.User
type memoryDataBasse struct {
	db userStore
}

var (
	mu sync.Mutex
)

// Get implements MemoryDb.
func (m memoryDataBasse) Get(username *string) (*models.User, error) {
	for _, user := range m.db {
		if user.Username == *username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

// GetAll implements MemoryDb.
func (m memoryDataBasse) GetAll() ([]*models.User, error) {
	fmt.Printf("m.db: %v\n", m.db)
	return m.db, nil
}

// Insert implements MemoryDb.
func (m *memoryDataBasse) Insert(user *models.User) *models.User {
	mu.Lock()
	m.db = append(m.db, user)
	fmt.Printf("m.db: %v\n", m.db)
	mu.Unlock()
	return user
}

type MemoryDb interface {
	Insert(user *models.User) *models.User
	Get(username *string) (*models.User, error)
	GetAll() ([]*models.User, error)
}

func New() MemoryDb {
	db := make([]*models.User, 0)
	return &memoryDataBasse{
		db: db,
	}
}
