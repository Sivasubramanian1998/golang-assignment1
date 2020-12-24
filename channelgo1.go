package main

import "fmt"

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		c <- i * i
	}
	close(c)
}

func main() {
	fmt.Println("main channel")
	c := make(chan int)
	go squares(c)
	for {
		val, ok := <-c
		if ok == false {
			fmt.Println(val, ok, "<--Loop Broke")
			break
		} else {
			fmt.println(val, ok)
		}
	}
	fmt.println("main Stopped")
}
