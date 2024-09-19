package model

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "1.grzh@yandex.ru",
		Password: "password",
	}
}
