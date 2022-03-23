package main

import "fmt"

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	a := i()
	a()
	aa := a()

	b := i()
	b()
	b()
	bb := b()

	fmt.Println("Oi", aa, bb)

}
