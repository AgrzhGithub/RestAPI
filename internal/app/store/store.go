package store

import (
	"database/sql"
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/lib/pq"
)

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	store := &Store{db: db}
	store.userRepository = &UserRepository{store: store}
	return store
}

func (s *Store) User() *UserRepository {
	if s.userRepository == nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}
