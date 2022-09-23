package main

import "fmt"

func main() {
greet("Sultan Aulia", "Karawang, West Java")
}

func greet(name, address string) {
	fmt.Println("Assalamualaikum my beloved, folks! my name is", name)
	fmt.Println("I live in", address)
}