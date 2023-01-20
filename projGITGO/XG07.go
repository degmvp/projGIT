package main

import (
	"fmt"
)

type SomDoPato interface {
	emiteSom()
}
type VooDoPato interface {
	voe()
}
type Pato struct {
	nome string
	SomDoPato
	VooDoPato
}

func (p Pato) saudacao() {
	fmt.Printf("Oi, Eu sou o %s, eu falo", p.nome)
	p.emiteSom()
	p.voe()
	fmt.Println("")
}

type Quieto struct{}

func (q Quieto) emiteSom() { fmt.Println("... Nada...") }

type QuaQua struct{}

func (q QuaQua) emiteSom() { fmt.Println("QuaQuá!") }

type Squeak struct{}

func (s Squeak) emiteSom() { fmt.Println("Squeak...") }

type Falador struct{}

func (f Falador) emiteSom() { fmt.Println("como um Humano.") }

type Voador struct{}

func (v Voador) voe() { fmt.Println("Estou voando!") }

type NaoVoa struct{}

func (nv NaoVoa) voe() { fmt.Println("Não posso voar...") }

type Latido struct{}

func (l Latido) latir() { fmt.Println("Au au!") }
func NewPatoMadeira() Pato {
	return Pato{"Pato de Madeira", new(Quieto), new(NaoVoa)}
}
func NewPatoBorracha() Pato {
	return Pato{VooDoPato: new(NaoVoa), nome: "Pato de Borracha", SomDoPato: new(Squeak)}
}
func NewPatoReal() Pato {
	return Pato{
		VooDoPato: new(Voador),
		SomDoPato: new(QuaQua),
		nome:      "Pato Real",
	}
}
func NewPatoDonald() Pato {
	d := new(Pato)
	d.nome = "Pato Donald"
	d.SomDoPato = new(Falador)
	// d.SomDoPato = new(Latido)
	d.VooDoPato = new(NaoVoa)
	return *d
}
func main() {
	patos := []Pato{NewPatoMadeira(), NewPatoBorracha(), NewPatoReal(), NewPatoDonald()}
	for _, pato := range patos {
		pato.saudacao()
	}
}
