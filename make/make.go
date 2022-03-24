package main

import "fmt"

func main() {
	meu_slice := make([]string, 3, 6)

	meu_slice[0] = "banana"
	meu_slice[1] = "maca"
	meu_slice[2] = "laranja"

	meu_slice = append(meu_slice, "uva")
	meu_slice = append(meu_slice, "manga")
	meu_slice = append(meu_slice, "abacate")

	fmt.Println(meu_slice)
	fmt.Println(len(meu_slice), cap(meu_slice))

	meu_slice = append(meu_slice, "limao")

	fmt.Println("")
	fmt.Println(meu_slice)
	fmt.Println(len(meu_slice), cap(meu_slice))
}
