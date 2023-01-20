// Go program to illustrate the  
// use for loop using channel
package main
  
import "fmt"
  
// Main function
func main() {
      
    // using channel
    chnl := make(chan int)
    go func(){
        chnl <- 100
        chnl <- 1000
       chnl <- 10000
       chnl <- 100000
       close(chnl)
    }()
    for i:= range chnl {
       fmt.Println(i) 
    }
    
}