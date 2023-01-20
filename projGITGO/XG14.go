// Go program to illustrate
// concept of variable
package main
import "fmt"
  
func main() {
  
// Using short variable declaration
myvar1 := 39 
myvar2 := "GeeksforGeeks" 
myvar3 := 34.67
  
// Display the value and type of the variables
fmt.Printf("The value of myvar1 is : %d\n", myvar1)
fmt.Printf("The type of myvar1 is : %T\n", myvar1)
  
fmt.Printf("\nThe value of myvar2 is : %s\n", myvar2)
fmt.Printf("The type of myvar2 is : %T\n", myvar2)
  
fmt.Printf("\nThe value of myvar3 is : %f\n", myvar3)
fmt.Printf("The type of myvar3 is : %T\n", myvar3)
}