// Go program to pass arguments 
// in the anonymous function
package main
  
import "fmt"
  
func main() {
      
    // Passing arguments in anonymous function
  func(ele string){
      fmt.Println(ele)
  }("GeeksforGeeks")
    
}