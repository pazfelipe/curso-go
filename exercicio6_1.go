package main

import "fmt"

func retornaInt() int {
	return 1
}

func retornIntEString() (int, string) {
	return 51, "ola"
}

func main() {
	n, str := retornIntEString()
	fmt.Println(retornaInt(), n, str)

}
