// Go program to illustrate the
// concept of init() function
  
// Declaration of the main package
package main
  
// Importing package
import "fmt"
  
// Multiple init() function
func init() {
    fmt.Println("Welcome to init() function")
}
  
func init() {
    fmt.Println("Hello! init() function")
}
  
// Main function
func main() {
    fmt.Println("Welcome to main() function")
}