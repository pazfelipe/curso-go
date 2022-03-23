package main

import "fmt"

func main() {
	_, errs := fmt.Println("Hello, World!", 45, "Felipe Paz")
	fmt.Println(errs)
}
