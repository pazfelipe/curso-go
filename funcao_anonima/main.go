package main

import "fmt"

func main() {
	x := func(value int) int {
		return value * 2
	}

	y := func(value int) func(int) int {
		return func(value2 int) int {
			return value * value2
		}
	}

	func(v string) {
		fmt.Println(v)
	}("Uma função anonima")

	fmt.Println("Oi", x(2), y(2)(3))

}
