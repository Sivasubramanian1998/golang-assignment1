package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("et$")

	fmt.Println(re.FindString("cricket"))
	fmt.Println(re.FindString("hacked"))
	fmt.Println(re.FindString("racket"))
}
