package main

import "fmt"

func main() {

	my_map := struct {
		atributos       map[string]string
		caracteristicas []string
	}{
		atributos: map[string]string{
			"cor":   "preto",
			"rodas": "liga-leve",
		},
		caracteristicas: []string{"novo", "caro"},
	}

	fmt.Println(my_map)

}
