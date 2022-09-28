package main

import (
	"errors"
	"fmt"

	"github.com/howeyc/gopass"
)




func main() {
	
	defer Defer()
	// var password string
	var username string
	fmt.Println("Username:")
	fmt.Scanln(&username) 
	fmt.Println("Password:")
	// fmt.Scanln(&password)
	password, _ := gopass.GetPasswdMasked()

	if valid, err := akunValidasi(username, password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}
	
}

func Defer() {
	if r := recover(); r != nil {
		fmt.Println("Ada yang error nih bg,", r)
	} else {
		fmt.Println("Lancar bang")
	}
}

func akunValidasi(username string, password []byte) (string, error) {
	p := string(password[:])
	if username != "bettercallaul" || p != "mantepbanget" {
		return "", errors.New("Ada yang salah tuh")
	}

	return "berhasil mashook!!", nil
}
