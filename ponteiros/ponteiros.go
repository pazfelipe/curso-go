package main

import "fmt"

func main() {
	x := 10 // atribuo o valor à variável x

	y := &x // referencio o endereço na memória da variável x

	z := &y // pego o valor da variável x que foi passada por referência para a variável y

	fmt.Println(&x, *y, z)

	*y = 20

	fmt.Println(&x, *y, z, x)

}
