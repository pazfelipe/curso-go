package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string `json:"first"`
	Age   int    `json:"age"`
}

func main() {
	user1 := User{"felipepaz", 25}
	user2 := User{"gabrielaguedes", 21}
	user3 := User{"oiutrousuario", 37}

	users := []User{user1, user2, user3}

	sb, err := json.Marshal(users)

	if err != nil {
		fmt.Println("Erro: ", err)
	}

	fmt.Println(string(sb))

}
