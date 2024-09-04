package store

import (
	"RestApi/internal/app/model"
	"github.com/pkg/errors"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {
	if r.store == nil {
		return nil, errors.New("store is nil")
	}
	if err := r.store.db.QueryRow("INSERT INTO users (email, encrypted_password) VALUES ($1,$2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id,email,encrypted_password FROM users WHERE email = $1",
		email,
	).Scan(
		&u.ID,
		&u.Email,
		&u.EncryptedPassword,
	); err != nil {
		return nil, err
	}
	return u, nil
}
