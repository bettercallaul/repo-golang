package main

import (
	"fmt"
)

func main () {
// for i := 0; i <= 3; i++ {
// 	fmt.Println("Angka", i)
// }

// var i = 1

// for i <= 10 {
// 	fmt.Println("angka", i)
// 	i++

// 	if i == 5 {
// 		break
// 	}
// }

// for i := 1; i <= 10; i++ {
// 	if (i % 2 == 1){
// 		continue
// 	}

// 	if i > 8 {
// 		break
// 	}
// 	fmt.Println("angka", i)
// }

// for i := 0; i < 5; i++ {
// 	for j := i; j < 5; j++ {
// 		fmt.Println(j, " ")
// 	}
// 	fmt.Println()
// }


outerLoop:
for i := 0; i < 5 ; i++ {
	fmt.Println("Perulanagan ke - ", i+1)
	for j := 0; j < 5; j++ {
		if i == 2 {
			break outerLoop
		}
		fmt.Println(j, " ")
	}
	fmt.Println()
}


}
