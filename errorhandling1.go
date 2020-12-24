package main

import (
	"fmt"
)

type myError struct{}

func (myErr *myError) Error() string {
	return "Some errors encountered!!"
}

func main() {

	myErr := &myError{}
	fmt.Println(myErr)
}
