package teststore_test

import (
	"testing"

	"github.com/shikidy/hh_sso_service/internal/domain/models"
	"github.com/shikidy/hh_sso_service/internal/store"
	"github.com/shikidy/hh_sso_service/internal/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	user := models.TestUser(t)
	s.User().Create(user)

	assert.NotEqual(t, user.ID, 0)
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.New()
	user := models.TestUser(t)

	s.User().Create(user)

	foundUser, err := s.User().Find(user.ID)
	assert.NoError(t, err)

	assert.Equal(t, foundUser, user)

	badUser, err := s.User().Find(123321)

	assert.Error(t, err, store.ErrRecordNotFound.Error())
	assert.Empty(t, badUser)

}
