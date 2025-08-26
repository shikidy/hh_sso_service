package sqlstore_test

import (
	"testing"

	"github.com/shikidy/hh_sso_service/internal/domain/models"
	"github.com/shikidy/hh_sso_service/internal/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")
	s := sqlstore.New(db)

	assert.NoError(t, s.User().Create(models.TestUser(t)))
}

func TestUserRepository_Find(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	tu := models.TestUser(t)

	s.User().Create(tu)
	assert.NotEqual(t, tu.ID, 0)
	_, err := s.User().Find(tu.ID)
	assert.NoError(t, err)

}
