package main

func sender(c chan int) {
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	close(c)
}

func main() {
	c := make(chan int, 3)

	go sender(c)

	fmt.printf("length of channel c is %v and capacity of channel is %v\n", len(c), cap(c))

	for val := range c {
		fmt.printf("length of channel c after value '%v' read is %v\n", len(c), cap(c))
	}
}
