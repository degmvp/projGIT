// Golang program to check if
// the slice is nil or not
package main
 
import "fmt"
 
func main() {
 
    // creating slices
    s1 := []int{12, 34, 56}
    var s2 []int
 
    // If you try to run this commented
    // code compiler will give an error
    /*s3:= []int{23, 45, 66}
      fmt.Println(s1==s3)
    */
 
    // Checking if the given slice is nil or not
    fmt.Println(s1 == nil)
    fmt.Println(s2 == nil)
}