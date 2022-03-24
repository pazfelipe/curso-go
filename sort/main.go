package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Oi")
	ss := []string{"banana", "abacaxi", "maça", "uva", "laranja"}
	si := []int{1, 25, 16, 38, 41, 02, 5}
	sort.Strings(ss)
	sort.Ints(si)

	fmt.Println(ss)
	fmt.Println(si)
}
