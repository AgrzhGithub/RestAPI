package store_test

import (
	"RestApi/internal/app/model"
	"RestApi/internal/app/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("users")

	s := store.New(db)
	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestFindByEmail(t *testing.T) {
	db, teardown := store.TestDB(t, databaseURL)
	defer teardown("users")

	s := store.New(db)
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
