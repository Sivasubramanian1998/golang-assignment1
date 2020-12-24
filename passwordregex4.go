package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("place")
	c := "play"
	fmt.Println(re.ReplaceAllString("replay string", c))
	fmt.Println(re.ReplaceAllString("no played", c))
	fmt.Println(re.ReplaceAllString("palace", c))
}
