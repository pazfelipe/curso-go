package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type sedan struct {
	veiculo
	modeloluxo bool
}

type caminhonete struct {
	veiculo
	quatroRodas bool
}

func main() {

	veiculo1 := sedan{veiculo{4, "preto"}, false}
	veiculo2 := caminhonete{veiculo{2, "cinza"}, true}
	fmt.Println(veiculo1)
	fmt.Println("")
	fmt.Println(veiculo2)
	fmt.Println("")
	fmt.Println(veiculo1.modeloluxo)
	fmt.Println(veiculo2.quatroRodas)

}
