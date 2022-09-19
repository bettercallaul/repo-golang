package main

import "fmt"

func main() {

	// prinln
	fmt.Println("Sultan", "Aulia", "Belajar", "Golang")
	fmt.Println("sultan", "Aulia")

	// print
	fmt.Print("Sultan Aulia\n")
	fmt.Print("Belajar", " Golang\n")
	fmt.Print("Aulia", " ", "Golang\n")

	// variabel

	var nama string = "Sultan Aulia"
	var umur int = 23

	fmt.Println("kenalin nama saya ", nama)
	fmt.Println("dan umur saya ", umur)


	// multiple variable
	var petugas1, petugas2, petugas3 string = "Dede Supriatna", "Dede Gusnul", "Lukman Bahtiar"
	
	fmt.Println(petugas1)
	fmt.Println(petugas2)
	fmt.Println(petugas3)

	var pustakawan1, pustakawan2, pustakawan3 int = 35, 39, 23
	fmt.Println("test data", pustakawan1)
	fmt.Println("test data", pustakawan2)
	fmt.Println("test data", pustakawan3)

	// underscore variabel
	var testVar string = "mantep"
	var satuNama, duanNama, Alamat = "Sultan", "Aulia", "Bojongsoang"
	_, _, _, _ = testVar, satuNama, duanNama, Alamat
	fmt.Println("test", satuNama, duanNama, Alamat)
}


