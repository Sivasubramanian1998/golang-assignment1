package main

import "fmt"

type address struct{
    name string
    city string
    pincode int
}

func main(){
    var a address
    fmt.Println(a)

    a1 := address{"Shiva" , "Chennai" , 600040}
    fmt.Println("Address1 :",  a1)

    a2 := address{name: "Bharath", city: "Chennai", pincode: 600040}
    fmt.Println("Address2 :",  a2)

    a3 := address{city: "Kerala"}
    fmt.Println("Address3 :",  a3)
}