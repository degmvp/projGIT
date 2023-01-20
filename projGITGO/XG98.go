// Go program to illustrate how to check
// whether the given slice of ints is in
// sorted form or not
package main
  
import (
    "fmt"
    "sort"
)
  
// Main function
func main() {
  
    // Creating and initializing slices
    // Using shorthand declaration
    scl1 := []int{100, 200, 300, 400, 500, 600, 700}
    scl2 := []int{-23, 567, -34, 67, 0, 12, -5}
  
    // Displaying slices
    fmt.Println("Slices:")
    fmt.Println("Slice 1: ", scl1)
    fmt.Println("Slice 2: ", scl2)
  
    // Checking the slice is in sorted form or not
    // Using IntsAreSorted function
    res1 := sort.IntsAreSorted(scl1)
    res2 := sort.IntsAreSorted(scl2)
  
    // Displaying the result
    fmt.Println("\nResult:")
    fmt.Println("Is Slice 1 is sorted?: ", res1)
    fmt.Println("Is Slice 2 is sorted?: ", res2)
}