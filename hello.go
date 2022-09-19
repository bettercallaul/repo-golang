package main

import "fmt"

func main() {
	//prinln basic
	fmt.Println("Assalamualikum")
	fmt.Println("Dunia Tipu-Tipu")

	fmt.Println("Asalamualaikum Dunia")
	fmt.Println("Assalamualaikum", "Dunia")

	fmt.Print("Mahal Kita\n")
	fmt.Print("Mahal", " Kita\n")
	fmt.Print("Mahal", " ", "Kita\n")

	fmt.Println("-----------------")

	//variable data type
	var name string
	var age int = 23

	name = "Sultan"
	age = 23

	fmt.Println("Kenalin nama saya ==>  ", name)
	fmt.Println("umur saya  ==> ", age)
	fmt.Println("-----------------")

	//variable with no data type
	var namee = "Telkomsel"
	var agee = 15

	fmt.Printf("%T,%T\n", namee, agee)
	fmt.Println("-----------------")

	//short declaration
	namaWatashi := "Orbit"
	umurWatashi := 8

	fmt.Printf("%T,%T\n", namaWatashi, umurWatashi)
	fmt.Println("-----------------")

	//multiple variable declarations
	var petugas1, petugas2, petugas3 string = "Dede Gusnul", "Saepulloh", "Anto Budianto"
	var first, second, third int
	first, second, third = 3, 7, 6

	fmt.Println(petugas1, petugas2, petugas3)
	fmt.Println(first, second, third)

	var nameee, ageee, addresss = "Danu", 59, "Jawa Barat"
	data1, data2, data3 := "Juned", 65, 4.7

	fmt.Println(nameee, ageee, addresss)
	fmt.Println(data1, data2, data3)
	fmt.Println("-----------------")

	//underscore variable
	//to cover up unused variable
	var firstVariable string
	var name1, age1, address1 = "Peeta", 15, "Bintaro Sektor 9"

	_, _, _, _ = firstVariable, name1, age1, address1
	fmt.Printf("test underscore variabel > %T\n", firstVariable)
	fmt.Printf("Kenalin gue %s Meelark, umur gue baru  %d tahun, dan sekarang tinggal di %s\n", name1, age1, address1)
	fmt.Println("-----------------")

	//Integer
	var num1 uint8 = 32
	var num2 int8 = -21
	fmt.Printf("tipe data num1: %T\n", num1)
	fmt.Printf("tipe data num2: %T\n", num2)
	fmt.Println("-----------------")

	//Float
	var decimal1 float32 = 3.63
	fmt.Printf("decimal number: %f\n", decimal1)
	fmt.Printf("decimal number: %.3f\n", decimal1)
	fmt.Println("-----------------")

	//boolean
	var bond bool = true
	fmt.Printf("dibolehin gak nih guys? %t\n", bond)
	fmt.Println("-----------------")

	//string
	var pesan = `Selamat datang di Indomaret. Silakan mau beli apa guys, asal jangan depo buat slot aja ya, daripada depo mending belajar Golang`
	fmt.Println(pesan)
	fmt.Println("-----------------")

	//constant
	const namaLengkap string = "Eka Kurniawan"
	fmt.Printf("Seperti Dendam, Rindu Harus Dibayar Tuntas. Karya %s\n", namaLengkap)
	fmt.Println("-----------------")

	//Operators
	var value = 2 + 2*3
	fmt.Println(value)

	var value2 = (2 + 2) * 3
	fmt.Println(value2)

	fmt.Println("-----------------")

	//Relational Operators

	var firstCondition bool = 2 < 3
	var secondCondition bool = "Reza" == "Rahadian"
	var thirdCondition bool = 10 != 2.3
	var forthCondition bool = 11 <= 11

	fmt.Println("first condition:", firstCondition)
	fmt.Println("second condition:", secondCondition)
	fmt.Println("third condition:", thirdCondition)
	fmt.Println("forth condition:", forthCondition)

	fmt.Println("-----------------")

	//Logical Operators
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("Wrong && Right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("Wrong || Right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)

}