// Go program to illustrate the
// use of arithmetic operators
package main
   
import "fmt"
   
func main() {
   p:= 34
   q:= 20

   a:= 10
   b:= 2
      
   // Addition
   result1:= p + q
   fmt.Printf("Result of p + q = %d", result1)
      
   // Subtraction
   result2:= p - q
   fmt.Printf("\nResult of p - q = %d", result2)
      
   // Multiplication
   result3:= p * q
   fmt.Printf("\nResult of p * q = %d", result3)
      
   // Division
   result4:= p / q
   fmt.Printf("\nResult of p / q = %d", result4)
      
   // Modulus
   result5:= p % q
   fmt.Printf("\nResult of p %% q = %d", result5)

   result6:= a % b
   fmt.Printf("\nResult of a %% b = %d", result6)

}