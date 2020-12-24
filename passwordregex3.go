package main
import (
  "fmt"
  "regexp"
)

func main() { 
  re := regexp.MustCompile("s([a-z]+)ar")
  fmt.Println(re.FindStringSubmatch("sarkar"))
  fmt.Println(re.FindStringSubmatch("cricket"))
}