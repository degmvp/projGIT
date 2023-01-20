// Go program to illustrate the
// concept of Expression switch
// statement
package main

import (
	"fmt"
)

func main() {

	var value int
	fmt.Scan(&value)

	// Switch statement without an
	// optional statement and
	// expression
	switch {
	case value == 1:
		fmt.Println("jan")
	case value == 2:
		fmt.Println("fev")
	case value == 3:
		fmt.Println("mar")

	case value == 4:
		fmt.Println("abr")
	case value == 5:
		fmt.Println("mai")
	case value == 6:
		fmt.Println("jun")

	case value == 7:
		fmt.Println("jul")
	case value == 8:
		fmt.Println("ago")
	case value == 9:
		fmt.Println("set")

	case value == 10:
		fmt.Println("out")
	case value == 11:
		fmt.Println("nov")
	case value == 12:
		fmt.Println("dez")

	default:
		fmt.Println("Invalid")
	}

}
