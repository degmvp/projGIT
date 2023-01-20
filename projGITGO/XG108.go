package main

import "fmt"

func main() {
	var num int = 0
	if num < 0 {
		fmt.Println(num, "é negativo")
	} else {
		fmt.Println(num, "é positivo")
	}

	var nums [5]int
	for i := 0; i <= 4; i++ {
		nums[i] = i + 10
	}
	for _, num := range nums {
		fmt.Println(num)
	}

}