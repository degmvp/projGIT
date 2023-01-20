// Go program to illustrate the concept
// of trim in the slice of bytes
package main
  
import (
    "bytes"
    "fmt"
)
  
func main() {
  
    // Creating and trimming
    // the slice of bytes
    // Using Trim function
    res1 := bytes.Trim([]byte("****Welcome to GeeksforGeeks****"), "*")
    res2 := bytes.Trim([]byte("!!!!Learning how to trim a slice of bytes@@@@"), "!@")
    res3 := bytes.Trim([]byte("^^Geek&&"), "$")
  
    // Display the results
    fmt.Printf("Final Slice:\n")
    fmt.Printf("\nSlice 1: %s", res1)
    fmt.Printf("\nSlice 2: %s", res2)
    fmt.Printf("\nSlice 3: %s", res3)
}