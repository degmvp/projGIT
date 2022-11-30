package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("data/config.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}
