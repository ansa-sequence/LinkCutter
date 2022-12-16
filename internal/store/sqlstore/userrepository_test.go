package sqlstore_test

import (
	"LinkCutter/internal/model"
	"LinkCutter/internal/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("information")

	s := sqlstore.New(db)
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("information")

	s := sqlstore.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	_ = s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_RemoveByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseUrl)
	defer teardown("information")

	s := sqlstore.New(db)
	email := "email"
	err := s.User().RemoveByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	err = s.User().Create(u)
	assert.NoError(t, err)

	err = s.User().RemoveByEmail(u.Email)
	assert.NoError(t, err)
}
