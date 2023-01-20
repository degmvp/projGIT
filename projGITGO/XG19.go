package main
 
import "fmt"
 
func main(){
    const A = "GFG"
    var B = "GeeksforGeeks"
     
    // Concat strings.
    var helloWorld = A+ " " + B
    helloWorld += "!"
    fmt.Println(helloWorld)
     
    // Compare strings.
    fmt.Println(A == "GFG")  
    fmt.Println(B < A)
}