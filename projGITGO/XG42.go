// Go program to illustrate the
// concept of the call by value
package main
  
import "fmt"
  
// function which swap values
func swap(a, b int)int{
 
    var o int
    o= a
    a=b
    b=o
    
   return o
}
  
// Main function
func main() {
 var p int = 10
 var q int = 20
  fmt.Printf("p = %d and q = %d", p, q)
  
 // call by values
 swap(p, q)
   fmt.Printf("\np = %d and q = %d",p, q)

 // call by values
 var x int = 100
 var y int = 205
 swap(x, y)
   fmt.Printf("\nx = %d and y = %d",x, y)
}