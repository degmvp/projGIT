// Golang program to illustrate how
// to create a slice using a slice
// literal
package main
 
import "fmt"
 
func main() {
 
    // Creating a slice
    // using the var keyword
    var my_slice_1 = []string{"Geeks", "for", "Geeks"}
 
    fmt.Println("My Slice 1:", my_slice_1)
 
    // Creating a slice
    //using shorthand declaration
    my_slice_2 := []int{12, 45, 67, 56, 43, 34, 45}
    fmt.Println("My Slice 2:", my_slice_2)
}