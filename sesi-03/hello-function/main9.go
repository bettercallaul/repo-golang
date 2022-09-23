package main

import "fmt"

// panindrome
func main() {
var isPalindrome = func (str string) bool {
	var temp string = ""
	for i := len(str) - 1; i >= 0; i++ {
		temp += string(byte(str[i]))
	}

	return temp == str
} ("level")

fmt.Println("mantep nih :", isPalindrome)
}

