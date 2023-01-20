// Go program to illustrate how to
// copy an array by reference
package main
  
import "fmt"
  
func main() {
  
    // Creating and initializing an array
    // Using shorthand declaration
    my_arr1 := [6]int{12, 456, 67, 65, 34, 34}
  
    // Copying the array into new variable
    // Here, the elements are passed by reference
    my_arr2 := &my_arr1
  
    fmt.Println("Array_1: ", my_arr1)
    fmt.Println("Array_2:", *my_arr2)
  
    my_arr1[5] = 1000000
  
    // Here, when we copy an array 
    // into another array by reference
    // then the changes made in original 
    // array will reflect in the
    // copy of that array
    fmt.Println("\nArray_1: ", my_arr1)
    fmt.Println("Array_2:", *my_arr2)
}