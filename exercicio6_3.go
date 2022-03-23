package main

import "fmt"

func comDefer() {
	fmt.Println("Vai ser executado com defer")
}

func main() {
	defer comDefer()
	fmt.Println("Oi")

}
