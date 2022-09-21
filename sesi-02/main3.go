package main

import "fmt"

func main() {
	// var numbers [4]int
	// numbers = [4] int{2, 3, 5, 8}

	// var strings = [3] string{"sultan", "andi", "Aulia"}
	// fmt.Printf("%#v\n", numbers)
	// fmt.Printf("%#v\n", strings)

	// var fruits = [4] string{"apple", "banana", "mangga", "kiwi"}
	// fruits[0] = "banana"
	// fruits[1] = "apple"
	// fruits[2] = "kiwi"
	// fruits[3] = "mangga"

	// fmt.Printf("%#v \n", fruits)

	// var fruits = [3] string{"apel", "pisang", "jagung"}

	// for i, v := range fruits {
	// 	fmt.Printf("index: %d, value: %s\n", i, v)
	// }

	// // ==============

	// fmt.Println(strings.Repeat("#", 25))

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("index: %d, value: %s\n", i, fruits[i])
	//}

	balances := [2][4]int{{5, 6, 7, 66}, {8, 9, 11, 23}}

	for _, arr := range balances {
		fmt.Printf("%d \n", arr)
		for _, value := range arr {
			fmt.Printf("%d \n", value)
		}
	}
}