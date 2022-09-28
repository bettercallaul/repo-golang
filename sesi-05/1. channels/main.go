package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Channels (Implementing channel)
	c := make(chan string)

	go introduce("Muhammad", c)

	go introduce("Sultan", c)

	go introduce("Aulia", c)
	
	go introduce("Rahmat", c)

	fmt.Println(strings.Repeat("=", 25))

	fmt.Println("Channels (Implementing channel)")
	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)
	
	msg4 := <-c
	fmt.Println(msg4)

	close(c)

	fmt.Println(strings.Repeat("=", 25))


	// Channels (Channel with anonymous function)
	fmt.Println("Channels (Channel with anonymous function)")
	c2 := make(chan string)
	student2 := []string{"Sultan", "Ahmad", "Aulia", "Rahmat"}

	for _, v := range student2 {
		go func(student2 string) {
			fmt.Println("Student", student2)
			result := fmt.Sprintf("Hai, nama saya %s", student2)
			c2 <- result
		}(v)
	}
	for i := 0; i <= 3; i++ {
		print(c2)
	}

	close(c2)
	fmt.Println(strings.Repeat("=", 25))


	// Channels (Channel directions)
	fmt.Println("Channels (Channel directions)")
	c3 := make(chan string)

	student3 := []string{"Ray", "Vierra", "Lakxmana", "Hidayat"}

	for _, v := range student3 {
		go introduce2(v, c3)
	}

	for i := 1; i <= 3; i++ {
		print2(c3)
	}

	close(c3)
	fmt.Println(strings.Repeat("=", 25))

	// UnBuffered Channel (Buffered vs unbuffered channel)
	fmt.Println("UnBuffered Channel (Buffered vs unbuffered channel)")
	c4 := make(chan int)

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c4)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	d := <-c4
	fmt.Println("main goroutine received data", d)

	close(c4)
	time.Sleep(time.Second)
	fmt.Println(strings.Repeat("=", 25))
	// Buffered Channel (Buffered vs unbuffered channel)
	fmt.Println("Buffered Channel (Buffered vs unbuffered channel)")
	c5 := make(chan int, 3)

	go func(c chan int, ) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine %d starts sending data into the channel\n", i)
			c <- 1
			fmt.Printf("func goroutine %d after sending data into the channel\n", i)
		}

		close(c)
	}(c5)

	fmt.Println("main goroutine sleeps 2 second")
	time.Sleep(time.Second * 4)
	
	for v := range c5 {
		
		fmt.Println("main goroutine received value from channel: ", v)
	}
	close(c5)
	fmt.Println("=====================")

	// Channel (Select)
	fmt.Println("Channel (Select)")
	c6 := make(chan string)
	c7 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)

		c6 <- "hello"
	}()

	go func() {
		time.Sleep(2 * time.Second)

		c7 <- "Hail Himmler"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c6:
			fmt.Println("Received", msg1)
		case msg2 := <-c7:
			fmt.Println("Received", msg2)
		}
	}
	fmt.Println(strings.Repeat("=", 25))


}

func print(c2 chan string) {
	fmt.Println(<-c2)
}

func introduce(student string, c chan string) {
	result := fmt.Sprintf("halo kenalin nama saya %s", student)

	c <- result
}

func print2(c4 chan string) {
	fmt.Println(<-c4)
}

func introduce2(student string, c chan string) {
	result := fmt.Sprintf("halo kenalin nama saya %s", student)

	c <- result
}