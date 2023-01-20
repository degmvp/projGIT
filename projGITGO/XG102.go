
package main

import (
	"bufio"
	f "fmt"
	"os"
	s "strings"
)

func main() {
	var nome, nomeAbreviado, sobreNome, primeiraLetra string
	var ultimoEspaco int
	f.Println("Digite o nome: ")
	scan := bufio.NewReader(os.Stdin)
	nome, _ = (scan.ReadString('\n'))
	nome = s.Trim(nome, " ")
	f.Println("Nome = ", nome, "*")

	if !s.Contains(nome, " ") { // Nome não possui sobrenome
		nomeAbreviado = string(nome)
	} else {
		primeiraLetra = string(nome[0])
		f.Println("Primeira letra = ", primeiraLetra)
		ultimoEspaco = s.LastIndex(nome, " ") + 1
		sobreNome = nome[ultimoEspaco:]
		f.Println("SobreNome = ", sobreNome)

		nomeAbreviado = primeiraLetra + "." + sobreNome
	}
	f.Println("Nome abreviado = ", nomeAbreviado)
}
