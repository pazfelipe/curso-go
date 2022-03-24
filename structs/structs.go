package main

import "fmt"

type client struct {
	nome      string
	sobrenome string
}

func main() {
	client1 := client{
		nome:      "Pedro",
		sobrenome: "Silva",
	}

	client2 := client{"Marcos", "Santos"}

	fmt.Println(client1)
	fmt.Println(client2)

}

// 2477707 23/03/2022 - rodobe - atendente Helen
