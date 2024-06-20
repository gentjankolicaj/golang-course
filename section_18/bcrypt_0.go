//go:build ignore

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

//Bycrypt samples:Module for password encryption

func main() {
	password := "Password123"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(hashedPassword)
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		fmt.Println("Passwords don't match")
	} else {
		fmt.Println("Passwords matched.")
	}
}
