package main 
  
import ( 
    "fmt"
    "time"
) 
  
func main() { 
  
    fmt.Println("Main function") 
  
    // Creating Anonymous Goroutine 
    go func() { 
  
        fmt.Println("Welcome") 
    }() 
  
    time.Sleep(1 * time.Second) 
    fmt.Println("GoodBye!!") 
} 