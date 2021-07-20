package models

import "golang.org/x/crypto/bcrypt"

func LoadSampleUser() *User {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpassword"), 8)
	return &User{
		Name:     "Sample User",
		Password: string(hashedPassword),
	}
}
