package sqlstore

import (
	"database/sql"

	"github.com/shikidy/hh_sso_service/internal/domain/models"
	"github.com/shikidy/hh_sso_service/internal/store"
)

type UserRepository struct {
	store *SQLStore
}

func (r *UserRepository) Create(user *models.User) error {
	err := r.store.db.QueryRow(
		"INSERT INTO users (email, username, encrypted_password) VALUES ($1, $2, $3) RETURNING id",
		user.Email,
		user.Username,
		user.PasswordHash,
	).Scan(&user.ID)
	return err
}

func (r *UserRepository) Find(id int) (*models.User, error) {
	u := &models.User{}

	if err := r.store.db.QueryRow(
		"SELECT id, username, email, created_at, updated_at FROM users WHERE id = $1",
		id,
	).Scan(
		&u.ID,
		&u.Username,
		&u.Email,
		&u.CreatedAt,
		&u.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}
