package main

import "fmt"

type Pessoa struct {
	nome string
}

func (p Pessoa) Gritar(grito string) {
	fmt.Println(p.nome, "gritou:", grito)
}

func main() {
	p1 := Pessoa{
		nome: "Felipe",
	}

	p1.Gritar("Oi")

}
