package users

import (
	"natasha/src/database"
)

func handleRegisterUser(user *User) error {
	return database.Repo.Create(user).Error
}
