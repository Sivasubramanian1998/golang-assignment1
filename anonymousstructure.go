package main

import "fmt"

func main(){
    details := struct{
    firstname string
    lastname string
    rollno int
    zipcode int
}{
    firstname : "Bharath",
    lastname : "Shiva",
    rollno : 21,
    zipcode : 600040,
}
fmt.Println(details)
}


