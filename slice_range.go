package main

func main() {
	slice := []string{"banana", "maca", "laranja"}

	for indice, valor := range slice {
		println(indice, valor)
	}
}
