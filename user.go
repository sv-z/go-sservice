package sservice

import "fmt"

type User struct {
	Id       int64
	Email    string `pg:",notnull,unique"`
	Password string `pg:",notnull"`
}

func (u User) String() string {
	return fmt.Sprintf("User<%d %s>", u.Id, u.Email)
}

// UserRepository provides access a user store.
type UserRepository interface {
	Store(user *User) error
	Delete(user *User) error
	GetById(userId int64) (*User, error)
	FindByEmail(email string) (*User, error)
	GetAll() ([]User, error)
}
