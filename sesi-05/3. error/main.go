package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("=", 25))
	// Error
	fmt.Println("Error")
	var number int
	var err error

	number, err = strconv.Atoi("123GH")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	number, err = strconv.Atoi("123")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}
	fmt.Println(strings.Repeat("=", 25))
	// Error (Custom error)
	fmt.Println("enter the password")
	var password string

	fmt.Scanln(&password)

	if valid, err := validPassword(password); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valid)
	}
	fmt.Println(strings.Repeat("=", 25))
}

// func validation password
func validPassword(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password harus memiliki lebih dari 4 kata")
	}

	return "Valid password", nil
}