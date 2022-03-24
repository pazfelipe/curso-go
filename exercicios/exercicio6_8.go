package main

import "fmt"

func retornaFuncao() func() {
	return func() {
		fmt.Println("Retorna uma função dentro de uma função")
	}
}

func main() {

	x := retornaFuncao()

	x()

}
