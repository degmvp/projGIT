// Go program to illustrate
// concept of variable
package main
import "fmt"
  
func main() {
  
// Using short variable declaration
// Multiple variables of same types
// are declared and initialized in 
// the single line
myvar1, myvar2, myvar3 := 800, 34, 56
  
// Display the value and 
// type of the variables
fmt.Printf("The value of myvar1 is : %d\n", myvar1)
fmt.Printf("The type of myvar1 is : %T\n", myvar1)
  
fmt.Printf("\nThe value of myvar2 is : %d\n", myvar2)
fmt.Printf("The type of myvar2 is : %T\n", myvar2)
  
fmt.Printf("\nThe value of myvar3 is : %d\n", myvar3)
fmt.Printf("The type of myvar3 is : %T\n", myvar3)
}