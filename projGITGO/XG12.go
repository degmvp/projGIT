// Go program to illustrate
// concept of variable
package main

import "fmt"

func main() {

	// Variable declared and
	// initialized without expression
	var myvariable1 int
	var myvariable2 string
	var myvariable3 float64

	// Display the zero-value of the variables
	fmt.Printf("The value of myvariable1 is : %d\n",
		myvariable1)

	fmt.Printf("The value of myvariable2 is : %s\n",
		myvariable2)

	fmt.Printf("The value of myvariable3 is : %f\n",
		myvariable3)

	fmt.Printf("The type of myvariable1 is : %T\n",
		myvariable1)

	fmt.Printf("The type of myvariable2 is : %T\n",
		myvariable2)

	fmt.Printf("The type of myvariable3 is : %T\n",
		myvariable3)
}
