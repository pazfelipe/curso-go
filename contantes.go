package main

import "fmt"

const (
	a = iota
)

func main() {
	const x = 10

	fmt.Printf("%T\n", x)
	fmt.Println(a)
}
