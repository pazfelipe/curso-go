package main

import "fmt"

func soma(arg int) int {
	return arg + 1
}

func naoRetornaNada() {
	fmt.Println("Não retorna nada")
}

func oi(nome string, args ...int) {
	fmt.Printf("%v, %v, %T\n", nome, args, args)
}

func somaMto(numeros ...int) int {
	soma := 0
	for _, value := range numeros {
		soma += value
	}

	return soma
}

func main() {
	fmt.Println(soma(1))
	naoRetornaNada()

	oi("Felipe", 1, 4, 3, 4, 5, 1234123)

	fmt.Println(somaMto(1, 23, 4, 234, 4, 5, 2, 43, 2))

}
