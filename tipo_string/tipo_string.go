package main

import "fmt"

func main() {
	s := "Hello, World!"
	sb := []byte(s)

	// fmt.Printf("%v, %T\n", s, s)
	// fmt.Printf("%v, %T", sb, sb)

	for _, v := range sb {
		fmt.Printf("%v, %T | %#U\n", v, v, v)
	}

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v, %T | %#U\n", s[i], s[i], s[i])
	}
}
