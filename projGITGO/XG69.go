// Go program to illustrate the
// concept of anonymous structure
package main
  
import "fmt"
  
// Main function
func main() {
  
    // Creating and initializing
    // the anonymous structure
    Element := struct {
        name      string
        branch    string
        language  string
        Particles int
    }{
        name:      "Pikachu",
        branch:    "ECE",
        language:  "C++",
        Particles: 498,
    }
  
    // Display the anonymous structure
    fmt.Println(Element)
}