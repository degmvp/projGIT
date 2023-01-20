
// Go program to illustrate how to create slices
// Using make function
package main
 
import "fmt"
 
func main() {
 
    // Creating an array of size 7
    // and slice this array  till 4
    // and return the reference of the slice
    // Using make function
    var my_slice_1 = make([]int, 4, 7)
    fmt.Printf("Slice 1 = %v, \nlength = %d, \ncapacity = %d\n",
                   my_slice_1, len(my_slice_1), cap(my_slice_1))
 
    // Creating another array of size 7
    // and return the reference of the slice
    // Using make function
    var my_slice_2 = make([]int, 7)
    fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n",
                   my_slice_2, len(my_slice_2), cap(my_slice_2))
     
}