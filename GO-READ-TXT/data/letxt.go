package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("data/1wrt.txt")
	if err != nil {
		panic(err)
		return
	}

	for {
		files, err := f.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("ERROR: %v\n", err)
			continue
		}
		fmt.Println(files[0].Name())
	}

}
