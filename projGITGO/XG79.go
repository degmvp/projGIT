// Go program to illustrate value type array
package main
 
import "fmt"
 
func main() {
     
// Creating an array whose size
// is represented by the ellipsis
my_array:= [...]int{100, 200, 300, 400, 500}
fmt.Println("Original array(Before):", my_array)
 
// Creating a new variable
// and initialize with my_array
new_array := my_array
 
fmt.Println("New array(before):", new_array)
 
// Change the value at index 0 to 500
new_array[0] = 500
 
fmt.Println("New array(After):", new_array)
 
fmt.Println("Original array(After):", my_array)
}