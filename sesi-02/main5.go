package main

import (
	"fmt"
)

func main () {
	var namaOrang map [string]string //declare
	namaOrang = map[string]string{} //initialize

	namaOrang["NAMA"] = "Lulu"
	namaOrang["UMUR"] = "22"

	namaOrang["Alamat"] = "Karawang Barat"

	fmt.Println("Nama: ", namaOrang["NAMA"])
	fmt.Println("Umur: ", namaOrang["UMUR"])
	fmt.Println("Alamat: ", namaOrang["Alamat"])


	
	var suckItSee = map [string]string {
		"Judul Film": "Bisikan Jenazah",
		"Tahun Rilis": "2021",
		"Nama Pemeran": "Aldi Taher",
	}

	value, exist := suckItSee["Tahun Rilis"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("gaada")
	}


	// looping with map (more simple way)
	for key, value := range suckItSee {
		fmt.Println(key, ":", value)
	}

	//delete value

	delete(suckItSee, "Tahun Rilis")
	value, exist = suckItSee["Tahun Rilis"] 
	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("ya ndak tau")
	} 
	fmt.Println("Abis Diapus: ", suckItSee)

// combining slice with map
// var ArcticMonkeys = []map[string]string {
// 	{"Name Song": "505", "Songwriter": "Alex Turner"},
// 	{"Name Song": "A Certain Romance", "Songwriter": "Arctic Monkeys"},
// 	{"Name Song": "Miracle Aligner", "Songwriter": "The Last Shadow Puppets"},
// }

// for i, ArcticMonkeys := range ArcticMonkeys {
// 	fmt.Printf("Index: %d, Name Song: %s, Songwriter: %s\n", i, ArcticMonkeys["Name Song"], ArcticMonkeys["Songwriter"])
// }


}