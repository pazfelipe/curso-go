package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("Oi")
	senha := "01091989"
	senhaerrada := "01091988"

	sb, err := bcrypt.GenerateFromPassword([]byte(senhaerrada), 17)

	if err != nil {
		fmt.Println("Erro: ", err)
	}

	fmt.Println(string(sb))

	err = bcrypt.CompareHashAndPassword(sb, []byte(senha))

	if err != nil {
		fmt.Println("Senha errada: ", err)
	}

}
