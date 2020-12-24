package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can't divide by zero")
	} else {
		return (a / b), nil
	}
}

func main() {
	if result, err := divide(4, 0); err != nil {
		fmt.Println("Error Occured :", err)
	} else {
		fmt.Println("4/0 is :", result)
	}
}
