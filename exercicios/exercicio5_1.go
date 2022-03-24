package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	pessoa1 := pessoa{
		nome:      "Felipe",
		sobrenome: "Paz",
		sabores:   []string{"branco", "abacaxi"},
	}

	for _, v := range pessoa1.sabores {
		fmt.Println("Felipe", v)
	}

	pessoa2 := pessoa{"Gabriela", "Guedes", []string{"branco", "laranja"}}

	fmt.Println("")

	for _, v := range pessoa2.sabores {
		fmt.Println("Gabriela:", v)
	}

}
