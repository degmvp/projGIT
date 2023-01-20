
package main

import "fmt"

// Main function
func main() {

	// calling the function
	// function returns two values which are
	// assigned to mul and blank identifier
	mul, _ := mul_div(105, 7)
	// only using the mul variable
	fmt.Println("105 x 7 = ", mul)

	mulx, _ := mul_div(30, 15)
	fmt.Println("30 * 15 = ", mulx)

}

// function returning two
// values of integer type
func mul_div(n1 int, n2 int) (int, int) {

	// returning the values
	return n1 * n2, n1 / n2

}
