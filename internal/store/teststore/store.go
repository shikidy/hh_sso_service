package teststore

import (
	"github.com/shikidy/hh_sso_service/internal/domain/models"
	"github.com/shikidy/hh_sso_service/internal/store"
)

type TestStore struct {
	userRepostiory *UserRepository
}

func New() *TestStore {
	return &TestStore{}
}

func (s *TestStore) User() store.UserRepository {
	if s.userRepostiory != nil {
		return s.userRepostiory
	}
	s.userRepostiory = &UserRepository{
		store: s,
		users: make(map[int]*models.User),
	}
	return s.userRepostiory
}
