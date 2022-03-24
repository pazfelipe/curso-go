package main

import "fmt"

type meu_tipo int

var x meu_tipo

func main() {
	fmt.Printf("%v\n", x)
	x = 42
	fmt.Printf("%v\n%T", x, x)
}
