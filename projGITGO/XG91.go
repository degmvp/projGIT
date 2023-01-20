// Golang program to illustrate the iterating over
// a slice using range in for loop without an index
package main
 
import "fmt"
 
func main() {
 
    // Creating a slice
    myslice := []string{"This", "is", "the",
        "tutorial", "of", "Go", "language"}
 
    // Iterate slice
    // using range in for loop
    // without index
    for _, ele := range myslice {
        fmt.Printf("Element = %s\n", ele)
    }
}