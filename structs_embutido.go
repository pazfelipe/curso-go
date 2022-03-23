package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario float64
}

func main() {
	pessoa1 := pessoa{
		nome:  "Felipe",
		idade: 23,
	}

	pessoa2 := pessoa{"Marcos", 34}

	profissional1 := profissional{
		pessoa: pessoa{
			nome:  "Outro",
			idade: 15,
		},
		titulo:  "Estudante",
		salario: 150.50,
	}

	profissional2 := profissional{pessoa{"João", 25}, "Programador", 1000.50}

	fmt.Println(pessoa1)
	fmt.Println(profissional1)

	fmt.Println("Caminho todo do profissional1: ", profissional1.pessoa.nome)

	// Caso não tenha conflitos de nomes, é possível acessar o atributo do scruct diretamente
	// sem precisar montar todo o caminho até o valor desejado
	fmt.Println("Caminho encurtado do profissional1: ", profissional1.nome)

	fmt.Println(profissional1.titulo)

	fmt.Println(pessoa2)
	fmt.Println(profissional2)

}
