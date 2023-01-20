
// Go program to illustrate the
// concept of ellipsis in an array
package main
 
import "fmt"
 
func main() {
     
// Creating an array whose size is determined
// by the number of elements present in it
// Using ellipsis
myarray:= [...]string{"GFG", "gfg", "geeks",
                    "GeeksforGeeks", "GEEK"}
 
fmt.Println("Elements of the array: ", myarray)
 
// Length of the array
// is determine by
// Using len() method
fmt.Println("Length of the array is:", len(myarray))
}