// Go program to illustrate the  
// use for loop using string
package main
  
import "fmt"
  
// Main function
func main() {
      
    // String as a range in the for loop
    for i, j:= range "XabCd" {
       fmt.Printf("The index number of %U is %d\n", j, i) 
    }
    
}