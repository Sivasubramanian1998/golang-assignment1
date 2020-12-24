package main

import(
    "errors"
    "fmt"
)

func main(){
    myErr := errors.New("Some errors occured!!")

    fmt.Println("Type of myErr is %T \n", myErr)
    fmt.Println("value of myErr is %#v \n", myErr)
}