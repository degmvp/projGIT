// Go program to illustrate
// use of an anonymous function
package main
  
import "fmt"
  
func main() {
      
    // Assigning an anonymous 
   // function to a variable
   value := func(){
      fmt.Println("Welcome! to GeeksforGeeks")
  }
  value()
    
}