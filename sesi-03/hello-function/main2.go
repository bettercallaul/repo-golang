package main

import (
	"fmt"
	"strings"
)

func main() {
var name = []string{"Alex","Turner"}
var printMsg = greet("Assalamualaikum ", name)
fmt.Println(printMsg)
}

func greet(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")
	var result string = fmt.Sprint("%s %s", msg, joinStr)
	return result
}