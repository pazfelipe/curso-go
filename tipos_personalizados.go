package main

import "fmt"

type dog int

var x dog

func main() {
	x = 10
	fmt.Println(x)

	b := fmt.Sprint(x)

	fmt.Printf("%T\n", b)
}
