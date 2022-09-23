package main

import (
	"fmt"
	"strings"
)

// panindrome
func main() {
var studentLists = []string{"Ariel", "niceso", "LLawliet", "Yagamoi"}
var find = findStudent(studentLists)
fmt.Println(find("Alex"))
}

func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s doesn't exist", s)
		}
		return fmt.Sprintf("we found %s at position %d", s, position)
	}
}

