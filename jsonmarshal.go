package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Sales     int    `json:"sales"`
	Age       int    `json:"age"`
	Developer string `json:"developer"`
}

func main() {

	author := Author{Sales: 3, Age: 25, Developer: "Stephen King"}
	book := Book{Title: "the shining", Author: author}

	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))
}
