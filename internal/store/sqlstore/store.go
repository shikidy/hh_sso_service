package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/shikidy/hh_sso_service/internal/store"
)

type SQLStore struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) store.Store {

	return &SQLStore{
		db: db,
	}
}

func (s *SQLStore) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}

	return s.userRepository
}
