package main

import "fmt"

func main() {
	words := []string{"Degmar", "Barbosa"}
	fmt.Println(Join(words))

	numbers := []int{5, 6}
	fmt.Println(JoinInts(numbers))
}

func Join(things []string) (result string) {
	for _, v := range things {
		result += v
	}
	return result
}

func JoinInts(things []int) (result string) {
	for _, v := range things {
		result += fmt.Sprint(v)
	}
	return result
}
