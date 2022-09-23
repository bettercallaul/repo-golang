package main

import "fmt"

func main() {
	var numbers = []int{2, 5, 8, 11, 15, 18, 21}
	var find = temukanNomorGanjil(numbers, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("jumlah nomor ganjil:",find)
}
func temukanNomorGanjil(numbers []int, callback func(int) bool) int {
	var jumlahNomorGanjil int
	for _, v := range numbers {
		if callback(v) {
			jumlahNomorGanjil++
		}
	}
	return jumlahNomorGanjil
}