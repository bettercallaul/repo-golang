package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("=", 25))
	// Recover

	defer catchErr()

	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
}

// catching error
func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Apllication running perfectly")
	}
}

// func validation password
func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("Password harus memiliki 4 karakter")
	}

	return "Valid password", nil
}