package main

import (
	"errors"
	"fmt"
)

func main() {
	myErr := errors.New("Some Errors Encountered!!")
	fmt.Println(myErr)
}
