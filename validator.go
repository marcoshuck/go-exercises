package main

import (
	"fmt"
	"strings"
	"unicode"
)

type user struct {
	ID       int
	Username string
	Password string
	Email    string
}
type validator interface {
	validatePassword()
	validateEmail()
}

//pwd should be 6 characters min, only ASCII characters.
func (u user) validatePassword(p string) bool {
	for len(p) >= 6 {
		for _, c := range p {
			if c > unicode.MaxASCII {
				break
			}
		}
		return true
	}
	return false
}

// email should have @ symbol and end with .com
func (u user) validateEmail(e string) bool {
	end := ".com"
	at := "@"
	u.Email = e
	if strings.Contains(e, at) {
		if strings.Contains(e, end) {
			return true
		}
	}
	return false
}

func main() {
	user1 := user{
		ID:       007,
		Username: "Santy",
		Password: "Pass12",
		Email:    "test@yaho.com",
	}
	fmt.Println("Username: ", user1.Username, "with ID: ", user1.ID, "email validation=", user1.validateEmail(user1.Email), "and password validation=", user1.validatePassword(user1.Password))
}
