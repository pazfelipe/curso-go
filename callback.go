package main

import "fmt"

// Recebe um slice de inteiros e retorna a soma de todos os elementos
// esse ...int é igual o rest operator do javascript
func soma(values ...int) int {
	total := 0

	for _, v := range values {
		total += v
	}

	return total
}

// Recebe uma função e um slice de inteiros e retorna a soma dos elementos do slice
// Vai ser parecido com o callback do javascript
func somentePares(f func(x ...int) int, y ...int) int {
	var slice []int

	for _, v := range y {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}

func somenteImpares(f func(x ...int) int, numeros ...int) (int, int) {
	var slice []int

	for _, numero := range numeros {
		if numero%2 != 0 {
			slice = append(slice, numero)
		}
	}

	total := f(slice...)
	return total, len(slice)
}

func main() {
	t := somentePares(soma, []int{1, 2, 3, 4, 5, 10}...)
	impares, totalImpares := somenteImpares(soma, []int{1, 3, 4, 5, 6, 7, 9, 10}...)
	fmt.Println(t)
	fmt.Println(impares, totalImpares)

}
