// Golang program to illustrate the
// iterating over a slice using
// for loop
package main
 
import "fmt"
 
func main() {
 
    // Creating a slice
    myslice := []string{"This", "is", "the", "tutorial",
        "of", "Go", "language"}
 
    // Iterate using for loop
    for e := 0; e < len(myslice); e++ {
        fmt.Println(myslice[e])
    }
}