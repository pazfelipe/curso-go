package main

import "fmt"

func main() {
	// O operador curto é usado para criar uma variável e atribuir um valor
	// Só é usado dentro de scoped blocks, ou seja, não é possível usar fora de um bloco

	x := 10
	y := "Felipe Paz"

	fmt.Printf("Value: %v, Type: %T\n", x, x)
	fmt.Printf("Value: %v, Type: %T\n", y, y)

	// X aqui não está sendo declarada novamente mas sendo atribuído um novo valor.
	// Como o operador curto necessita que tenha, no mínimo, uma variável a ser declarada
	// nesse caso, o z está sendo declarado e atribuído um valor, enquanto que x
	// está recebendo um novo valor.
	x, z := 10, 20
	fmt.Println(x, z)
}
