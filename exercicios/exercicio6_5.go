package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

type circulo struct {
	raio float64
}

type info interface {
	area()
}

func medida(i info) {
	i.area()
}

// Comentário aqui
func (q quadrado) area() {
	resultado := q.lado * q.lado
	fmt.Println("A area do quadrado é:", resultado)
}

func (c circulo) area() {
	resultado := math.Pi * 2 * c.raio
	fmt.Println("A area do círculo é", resultado)
}

func main() {
	q := quadrado{5}
	c := circulo{10}

	// q.area()
	// c.area()

	medida(q)
	medida(c)
}
