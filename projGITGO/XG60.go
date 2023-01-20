// Go program to illustrate pointer receiver
package main
 
import "fmt"
 
// Author structure
type author struct {
    name      string
    branch    string
    particles int
}
 
// Method with a receiver of author type
func (a *author) show(abranch string) {
    (*a).branch = abranch
}
 
// Main function
func main() {
 
    // Initializing the values
    // of the author structure
    res := author{
        name:   "Sona",
        branch: "CSE",
    }
 
    fmt.Println("Author's name: ", res.name)
    fmt.Println("Branch Name(Before): ", res.branch)
 
    // Creating a pointer
    p := &res
 
    // Calling the show method
    p.show("ECE")
    fmt.Println("Author's name: ", res.name)
    fmt.Println("Branch Name(After): ", res.branch)
}