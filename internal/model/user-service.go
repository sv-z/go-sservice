package model

import (
	"golang.org/x/crypto/bcrypt"
)

// Crate new user
func CreateNewUser(email string, rawPassword string) *User {
	return &User{
		Email:             email,
		EncryptedPassword: encryptString(rawPassword),
	}
}

func ComparePassword(user *User, rawPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(rawPassword)) != nil
}

// Generate md5 hash ...
func encryptString(text string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return string(b)
}
