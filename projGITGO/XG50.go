// Go program to illustrate
// use of anonymous function
package main
  
import "fmt"
  
 // Returning anonymous function 
 func GFG() func(i, j string) string{
     myf := func(i, j string)string{
          return i + j + "GeeksforGeeks"
     }
    return myf
 }
    
func main() {
    value := GFG()
    fmt.Println(value("Welcome ", "to "))
}