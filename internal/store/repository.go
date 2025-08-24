package store

import "github.com/shikidy/hh_sso_service/internal/domain/models"

type UserRepository interface {
	Create(*models.User)
	Find(int) (*models.User, error)
}
