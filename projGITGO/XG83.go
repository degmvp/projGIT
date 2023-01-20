
// Go program to illustrate how to pass an
// array as an argument in the function
package main
 
import "fmt"
 
// This function accept
// an array as an argument
func myfun(a [6]int, size int) int {
    var k, val, r int
 
    for k = 0; k < size; k++ {
        val += a[k]
    }
 
    r = val / size
    return r
}
 
// Main function
func main() {
 
    // Creating and initializing an array
    var arr = [6]int{67, 59, 29, 35, 4, 34}
    var res int
 
    // Passing an array as an argument
    res = myfun(arr, 6)
    fmt.Printf("Final result is: %d ", res)
}