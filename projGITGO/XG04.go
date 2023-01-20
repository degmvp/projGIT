package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	m["k3"] = 45
	m["k4"] = 83

	fmt.Println("map:", m)
	//v1------------------------------
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))
	//v2------------------------------
	v2 := m["k2"]
	fmt.Println("v2: ", v2)

	fmt.Println("len:", len(m))
	//v3--------------------------------
	v3 := m["k3"]
	fmt.Println("v3: ", v3)

	fmt.Println("len:", len(m))
	//v4--------------------------------
	v4 := m["k4"]
	fmt.Println("v4: ", v4)

	fmt.Println("len:", len(m))
	//deleta elemento k2
	//-----------------------------------
	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
