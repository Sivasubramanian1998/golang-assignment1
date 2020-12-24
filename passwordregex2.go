package main
import (
  "fmt"
  "regexp"
)

func main() {  
  re := regexp.MustCompile("tel")
  fmt.Println(re.FindStringIndex("telephone"))
  fmt.Println(re.FindStringIndex("carpet"))
  fmt.Println(re.FindStringIndex("cartel"))
}