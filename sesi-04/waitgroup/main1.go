package main

import (
	"fmt"
	"sync"
)

// waitgroup
func main () {
	fruits := []string{"doctor", "uber", "colorful", "alright", "really care"}
	var wg sync.WaitGroup
	for index, fruit := range fruits {
		wg.Add(1)
		go printFruit(index, fruit, &wg)
	}
	wg.Wait()

}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index => %d, fruit => %s\n", index, fruit)
	wg.Done()
}