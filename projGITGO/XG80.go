// Go program to illustrate
// how to compare two arrays
package main
 
import "fmt"
 
func main() {
 
// Arrays   
arr1:= [3]int{9,7,6}
arr2:= [...]int{9,7,6}
arr3:= [3]int{9,5,3}
 
// Comparing arrays using == operator
fmt.Println(arr1==arr2)
fmt.Println(arr2==arr3)
fmt.Println(arr1==arr3)
 
// This will give and error because the
// type of arr1 and arr4 is a mismatch
/*
arr4:= [4]int{9,7,6}
fmt.Println(arr1==arr4)
*/
}