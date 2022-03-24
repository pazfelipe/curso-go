package main

import "fmt"

func main() {
	for n := 65; n <= 90; n++ {
		fmt.Printf("%v\n", n)
		for i := 0; i <= 2; i++ {
			fmt.Printf("\t%#U\n", n)
		}
		fmt.Print("\n")
	}
}
