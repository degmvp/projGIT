package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("data/data1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	for i := 0; i < 10; i++ {

		num := i
		email := "wrt.txt"
		conc := fmt.Sprintf("%d%s", num, email)

		fmt.Println(conc)

		err = os.WriteFile(conc, []byte(conc), 0755)
		if err != nil {
			panic(err)
		}
	}

}
