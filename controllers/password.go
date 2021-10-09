package controllers

import (
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	return string(hash)
}
