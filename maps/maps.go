package main

import "fmt"

func main() {
	amigos := map[string]int{
		"Pedro":  23,
		"Carlos": 34,
	}

	valorSeExiste, ok := amigos["Pedro"]

	fmt.Println(valorSeExiste, ok)

	fmt.Println(amigos)

	for key, value := range amigos {
		fmt.Printf("Key: %v, Value: %d\n", key, value)
	}

	delete(amigos, "Pedro")

	fmt.Println(amigos)

}
