
// Go program to illustrate the concept
// of trim in the slice of bytes
package main
  
import (
    "bytes"
    "fmt"
)
  
func main() {
  
    // Creating and initializing
    // the slice of bytes
    // Using shorthand declaration
    slice_1 := []byte{'!', '!', 'G', 'e', 'e', 'k', 's', 'f',  
                 'o', 'r', 'G', 'e', 'e', 'k', 's', '#', '#'}
      
    slice_2 := []byte{'*', '*', 'A', 'p', 'p', 'l', 'e', '^', '^'}
      
    slice_3 := []byte{'%', 'g', 'e', 'e', 'k', 's', '%'}
  
    // Displaying slices
    fmt.Println("Original Slice:")
    fmt.Printf("Slice 1: %s", slice_1)
    fmt.Printf("\nSlice 2: %s", slice_2)
    fmt.Printf("\nSlice 3: %s", slice_3)
  
    // Trimming specified leading 
    // and trailing Unicodes points
    // from the given slice of bytes
    // Using Trim function
    res1 := bytes.Trim(slice_1, "!#")
    res2 := bytes.Trim(slice_2, "*^")
    res3 := bytes.Trim(slice_3, "@")
  
    // Display the results
    fmt.Printf("New Slice:\n")
    fmt.Printf("\nSlice 1: %s", res1)
    fmt.Printf("\nSlice 2: %s", res2)
    fmt.Printf("\nSlice 3: %s", res3)
  
}