package main

import "fmt"


func main () {

	// combining slice with map

var ArcticMonkeys = []map[string]string {
	{"Name Song": "505", "Songwriter": "Alex Turner"},
	{"Name Song": "A Certain Romance", "Songwriter": "Arctic Monkeys"},
	{"Name Song": "Miracle Aligner", "Songwriter": "The Last Shadow Puppets"},
}

for i, ArcticMonkeys := range ArcticMonkeys {
	fmt.Printf("Index: %d, Name Song: %s, Songwriter: %s\n", i, ArcticMonkeys["Name Song"], ArcticMonkeys["Songwriter"])
}
type second = uint
var jam second = 3600
fmt.Printf("tipe jam: %T\n", jam)


// looping over string (byte by byte)
album := "Suck It & See"
for i := 0; i < len(album); i++ {
	fmt.Printf("Index: %d, byte: %d\n", i, album[i])
}
}