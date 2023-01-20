/*
Algoritmos em GOLANG

Enunciado: Leia duas notas notas e calcule a média aritmética
*/
package main

import "fmt"

func main() {
	var nota1, nota2, media float64
	fmt.Print("Informe a nota 1: ")
	fmt.Scan(&nota1)
	fmt.Print("Informe a nota 2: ")
	fmt.Scan(&nota2)
	media = (nota1 + nota2) / 2.0
	fmt.Printf("A média entre %2.f e %2.f é %2.f", nota1, nota2, media)
}
