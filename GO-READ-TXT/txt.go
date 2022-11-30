package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("data/gene.go")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	err = os.WriteFile("data/wrtsaida.txt", []byte(data), 0755)
	if err != nil {
		panic(err)
	}
}
