package main

import "fmt"

// Defer é usado para executar uma expressão ao final do bloco de instrução
// Um exemplo de uso pode ser a escrita e leitura de um arquivo
// func abrirAquivo()
// defer func fecharAquivo()
// ... demais execuções dentro do bloco

func main() {
	defer fmt.Println("Usando defer: Veio primeiro")
	fmt.Println("Sem defer: Veio depois")

	fmt.Println("")

	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)

}
