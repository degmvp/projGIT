package main

//range no map itera com os pares chave/valor

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	soma := 0
	for _, num := range nums {
		soma += num
	}
	fmt.Println("soma:", soma)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	cvs := map[string]string{"a": "maçã", "b": "banana"}
	for c, v := range cvs {
		fmt.Printf("%s -> %s\n", c, v)
	}

	for i, c := range "DEG" {
		fmt.Println(i, c)
	}
}
