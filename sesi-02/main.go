package main

import "fmt"


func main () {
	var score = 10

	if (score > 7) {
		switch score {
		case 10:
			fmt.Println("Perfecto!")
		default:
			fmt.Println("Nice!")
			
		}
	} else {
		if (score == 5) {

		}else if (score == 3) {

		} else {
			fmt.Println("you can do it")
			if (score == 0) {
				fmt.Println("try harder")
			}
		}
	}

	switch  {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score >= 3): 
		fmt.Println("awesome")
	case score <5:
		fmt.Println("it is oke, but please harder")
	default:
		fmt.Println("not bad")
		fmt.Println("keren")
	}



		

}