package main

import "fmt"

func main() {
	x := 3

	switch x {
	case 34:
		fmt.Println("x is 10")
		fallthrough // isso significa que ele vai executar o case abaixo
	case 20:
		fmt.Println("x is 20 | Executado por causa do fallthrough")
	default:
		fmt.Println("x is not 10 or 20")
	}

	switch {
	case x == 34:
		fmt.Println("x is equal 34")
	case x == 20:
		fmt.Println("x is 20")
	default:
		fmt.Println("x is not 10 or 20")
	}

	switch x {
	case 1, 2:
		fmt.Println("x is 1 or 2")
	case 3, 4:
		fmt.Println("x is 3 or 4")
	default:
		fmt.Println("x is not 1 or 2 or 3 or 4")
	}
}
