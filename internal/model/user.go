package model

import (
	"fmt"
)

// UserRepository provides access a user store.
type UserRepository interface {
	// Save entity in db
	Save(user *User) error

	// Del entity from db
	Delete(user *User)

	// Get entity by id, if entity not found return error
	GetById(userId int) *User

	// Find one entity by email, if entity not found return nil
	FindByEmail(email string) *User

	// Find all entities by params
	GetAll() []*User
}

type User struct {
	Id                int    `json:"id"`
	Email             string `json:"email"`
	EncryptedPassword string `json:"-"`
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s>", u.Id, u.Email)
}
