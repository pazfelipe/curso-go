package main

import "fmt"

type meu_tipo int

var x meu_tipo
var y int

func main() {
	fmt.Printf("%v\n", x)
	x = 42
	y = int(x)
	fmt.Printf("%v\n%T\n", x, x)
	fmt.Printf("%v\n%T\n", y, y)
}
