package main

import "fmt"

func main() {
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Felipe",
		idade: 32,
	}

	fmt.Println(x)

}
