package main

import "fmt"

type Pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p Pessoa) NomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func main() {
	p1 := Pessoa{
		nome:      "Felipe",
		sobrenome: "Paz",
	}
	fmt.Println("Oi,", p1.NomeCompleto())

}
