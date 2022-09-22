package main

import (
	"fmt"
)

func main() {

// slice make function

// var fruits = make ([]string, 3)

// _ = fruits
// append
// fruits[0] = "gagah"
// fruits[1] = "jimi"
// fruits[2] = "labrador"

// fruits = append(fruits, "semakin", "memanas", "gugur")

// fmt.Printf("%#s", fruits)

// append with ellipsis

var fruits1 = []string{"racun", "dunia", "the changcuters", "hijrah"}
var fruits2 = []string{"labrador", "mastaka", "the panturas", "london"}

fruits1 = append(fruits1, fruits2...)

fmt.Printf("%#v\n", fruits1)

// copy function

nn := copy(fruits1, fruits2)

fmt.Println("Buah-buahan =>", fruits1)
fmt.Println("Buah-buahan2 =>", fruits2)
fmt.Println("penyalin cahaya =>", nn)

var karang = fruits1[3:4]
fmt.Printf("%#v\n", karang)

var sekaliLagi = []string{"maafkanlah", "ku tak bisa", "tinggalkan dirinya"}
sekaliLagi = append(sekaliLagi[0:2], "i love u so") 
fmt.Println(sekaliLagi)


	// mobil := []string{"Kodok", "Jimli", "X-Ranger", "Pagani"}
	// baruMobil := []string{}

	// baruMobil = append(baruMobil, mobil[0:2]...)

	// mobil[0] = "Tesla"
	// fmt.Println("nama mobil:", mobil)
	// fmt.Println("nama mobil baru:", baruMobil)
}