// Const demonstration using go.
package main
 
import (
    "fmt"
    "math"
)
 
const s string = "GeeksForGeeks"
 
func main() {
    fmt.Println(s)
 
    const n = 5
 
    const d = 3e10 / n
    fmt.Println(d)
 
    fmt.Println(int64(d))
 
    fmt.Println(math.Sin(n))
}