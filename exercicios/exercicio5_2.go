package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	my_map := make(map[string]pessoa)

	my_map["pessoa1"] = pessoa{
		nome:      "Felipe",
		sobrenome: "Paz",
		sabores:   []string{"branco", "abacaxi"},
	}

	my_map2 := map[string]pessoa{
		"pessoa1": {
			nome:      "Felipe",
			sobrenome: "Paz",
			sabores:   []string{"roxo", "azul"},
		},
		"pessoa2": {"Gabriela", "Guedes", []string{"rosa", "branco"}},
	}

	for key, value := range my_map {
		fmt.Println(key, ":", value)
	}

	fmt.Println("")

	for key, value := range my_map2 {
		fmt.Printf("\nMeu nome é: %v e gosto de:\n", key)

		for _, v := range value.sabores {
			fmt.Printf("- %v\n", v)
		}
	}

}
