// Go program to illustrate how the
// method can accept pointer and value
 
package main
 
import "fmt"
 
// Author structure
type author struct {
    name   string
    branch string
}
 
// Method with a pointer
// receiver of author type
func (a *author) show_1(abranch string) {
    (*a).branch = abranch
}
 
// Method with a value
// receiver of author type
func (a author) show_2() {
 
    a.name = "Gourav"
    fmt.Println("Author's name(Before) : ", a.name)
}
 
// Main function
func main() {
 
    // Initializing the values
    // of the author structure
    res := author{
        name:   "Sona",
        branch: "CSE",
    }
 
    fmt.Println("Branch Name(Before): ", res.branch)
 
    // Calling the show_1 method
    // (pointer method) with value
    res.show_1("ECE")
    fmt.Println("Branch Name(After): ", res.branch)
 
    // Calling the show_2 method
    // (value method) with a pointer
    (&res).show_2()
    fmt.Println("Author's name(After): ", res.name)
}
