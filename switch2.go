package main

import "fmt"

var x interface{}

func main() {
	x = 10.4234

	switch x.(type) {
	case int:
		fmt.Println("x is an int")
	case string:
		fmt.Println("x is a string")
	case bool:
		fmt.Println("x is a bool")
	case float64, float32:
		fmt.Println("x is a float")
	}
}
