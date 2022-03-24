package main

// Para criar um sort customizado,
// é preciso criar uma interface e implementar três métodos obrigatórios:
// Len, Less e Swap.

import (
	"fmt"
	"sort"
)

type Carro struct {
	Nome     string `json:"nome"`
	Potencia int    `json:"potencia"`
	Consumo  int    `json:"consumo"`
	Valor    int    `json:"valor"`
}

type OrderPorLucro []Carro
type OrderPorConsumo []Carro
type OrderPorPotencia []Carro

func (c OrderPorPotencia) Len() int           { return len(c) }
func (c OrderPorPotencia) Less(i, j int) bool { return c[i].Potencia < c[j].Potencia }
func (c OrderPorPotencia) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c OrderPorConsumo) Len() int           { return len(c) }
func (c OrderPorConsumo) Less(i, j int) bool { return c[i].Consumo < c[j].Consumo }
func (c OrderPorConsumo) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (c OrderPorLucro) Len() int           { return len(c) }
func (c OrderPorLucro) Less(i, j int) bool { return c[i].Valor > c[j].Valor }
func (c OrderPorLucro) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func main() {
	carros := []Carro{
		{"Fusca", 5, 10, 1500},
		{"Gol", 15, 9, 18000},
		{"Fox", 14, 12, 22500},
	}

	fmt.Printf("\nOrdenando crescentemente por potência...\n")
	sort.Sort(OrderPorPotencia(carros))
	fmt.Println(carros)

	fmt.Printf("\nOrdenando crescentemente por consumo...\n")
	sort.Sort(OrderPorConsumo(carros))
	fmt.Println(carros)

	fmt.Printf("\nOrdenando decrescentemente por lucro...\n")
	sort.Sort(OrderPorLucro(carros))
	fmt.Println(carros)
}
