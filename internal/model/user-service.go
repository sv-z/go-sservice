package model

import (
	"crypto/md5"
	"encoding/hex"
)

// Crate new user
func CreateNewUser(email string, rawPassword string) *User {
	return &User{
		Email:             email,
		EncryptedPassword: md5hash(rawPassword),
	}
}

// Generate md5 hash ...
func md5hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
