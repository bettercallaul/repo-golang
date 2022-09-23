package main

import (
	"fmt"
	"strings"
)

func main() {
profile("Alexx", "Mie Ayam", "Fried Chicken")
}

func profile(name string, makananFavorit ...string) {
	mergemakananFavorit := strings.Join(makananFavorit, ", ")
	fmt.Println("Oh Hi there, kenalin gue", name)
	fmt.Println("dan gue sering makan ", mergemakananFavorit)
}