// Golang program to illustrate the iterating
// over a slice using range in for loop
package main
 
import "fmt"
 
func main() {
 
    // Creating a slice
    myslice := []string{"This", "is", "the", "tutorial",
                                 "of", "Go", "language"}
 
    // Iterate slice
    // using range in for loop
    for index, ele := range myslice {
        fmt.Printf("Index = %d and element = %s\n", index+3, ele)
    }
}