package main

import "fmt"

func sending(s chan<- string) {
	s <- "good morning"
}

func main() {

	// Creating a bidirectional channel
	mychanl := make(chan string)

	// Here, sending() function convert
	// the bidirectional channel
	// into send only channel
	go sending(mychanl)

	// Here, the channel is sent
	// only inside the goroutine
	// outside the goroutine the
	// channel is bidirectional
	// So, it print good morning
	fmt.Println(<-mychanl)
}
