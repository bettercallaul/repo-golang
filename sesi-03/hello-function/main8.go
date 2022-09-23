package main

import "fmt"

func main() {
	var nomorGenap = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result

	}

	var numbers = []int{5, 7, 2, 10, 28, 73, 64, 21, 7}
	fmt.Println(nomorGenap(numbers...))
}
