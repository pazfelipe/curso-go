package main

import "fmt"

func funcao(x ...int) int {
	soma := 0
	for _, numero := range x {
		soma += numero
	}

	return soma
}

func funcao2(x []int) int {
	soma := 0
	for _, numero := range x {
		soma += numero
	}

	return soma
}

func main() {
	t := funcao([]int{1, 2, 3, 4, 5, 6, 6}...)
	t2 := funcao2([]int{1, 34, 5, 534, 56})
	fmt.Println(t)
	fmt.Println(t2)

}
