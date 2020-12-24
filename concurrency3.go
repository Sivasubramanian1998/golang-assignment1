package main 
  
import ( 
    "fmt"
    "time"
) 
  
// For goroutine 1 
func name() { 
  
    arr1 := [4]string{"shiva", "akash", "balaji", "parthiban"} 
  
    for t1 := 0; t1 <= 3; t1++ { 
      
        time.Sleep(150 * time.Millisecond) 
        fmt.Printf("%s\n", arr1[t1]) 
    } 
} 
  
// For goroutine 2 
func id() { 
  
    arr2 := [4]int{51, 4, 12, 36} 
      
    for t2 := 0; t2 <= 3; t2++ { 
      
        time.Sleep(500 * time.Millisecond) 
        fmt.Printf("%d\n", arr2[t2]) 
    } 
} 
  
// Main function 
func main() { 
    // calling Goroutine 1 
    go name() 
  
    // calling Goroutine 2 
    go id() 
  
    time.Sleep(3500 * time.Millisecond) 
} 