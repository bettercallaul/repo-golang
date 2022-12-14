package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// defer #1
	defer fmt.Println("defer function starts to execute")
	
	fmt.Println("Hai everyone")
	fmt.Println("Mantep Abiezz")
	fmt.Println("Welcome back to Go learning center")
	fmt.Println(strings.Repeat("=", 25))

	// defer #2
	callDeferFunc()
	fmt.Println("Hai everyone")
	fmt.Println(strings.Repeat("=", 25))

	mainn()

}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}

func mainn() {
		// Exit
		defer fmt.Println("Invoke with defer")

		fmt.Println("Before exiting")
		
		os.Exit(1)
	
		// exit tidak terlihat
		fmt.Println("Not Seen")
}