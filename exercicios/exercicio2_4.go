package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("%d, %#x, %b\n", x, x, x)
	y := x << 1
	fmt.Printf("%d, %#x, %b\n", y, y, y)
}
