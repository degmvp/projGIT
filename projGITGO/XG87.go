// Golang program to illustrate how to
// create slices from the slice
package main
 
import "fmt"
 
func main() {
 
    // Creating s slice
    oRignAl_slice := []int{90, 60, 40, 50,
        34, 49, 30}
 
    // Creating slices from the given slice
    var my_slice_1 = oRignAl_slice[1:5]
    my_slice_2 := oRignAl_slice[0:]
    my_slice_3 := oRignAl_slice[:6]
    my_slice_4 := oRignAl_slice[:]
    my_slice_5 := my_slice_3[2:4]
 
    // Display the result
    fmt.Println("Original Slice:", oRignAl_slice)
    fmt.Println("New Slice 1:", my_slice_1)
    fmt.Println("New Slice 2:", my_slice_2)
    fmt.Println("New Slice 3:", my_slice_3)
    fmt.Println("New Slice 4:", my_slice_4)
    fmt.Println("New Slice 5:", my_slice_5)
}