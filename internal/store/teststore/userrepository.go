package teststore

import (
	"github.com/shikidy/hh_sso_service/internal/domain/models"
	"github.com/shikidy/hh_sso_service/internal/store"
)

type UserRepository struct {
	store *TestStore
	users map[int]*models.User
}

func (r *UserRepository) Create(user *models.User) {
	newId := len(r.users) + 1
	user.ID = newId
	r.users[newId] = user
}

func (r *UserRepository) Find(id int) (*models.User, error) {
	value, ok := r.users[id]
	if !ok {
		return nil, store.ErrRecordNotFound
	}
	return value, nil
}
