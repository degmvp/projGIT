
// Go program to illustrate the
// method with struct type receiver
package main
 
import "fmt"
 
// Author structure
type author struct {
    name      string
    branch    string
    particles int
    salary    int
}
 
// Method with a receiver
// of author type
func (a author) show() {
 
    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Branch Name: ", a.branch)
    fmt.Println("Published articles: ", a.particles)
    fmt.Println("Salary: ", a.salary)
}
 
// Main function
func main() {
 
    // Initializing the values
    // of the author structure
    res := author{
        name:      "Deg",
        branch:    "JP",
        particles: 77,
        salary:    3400,
    }
 
    // Calling the method
    res.show()
}