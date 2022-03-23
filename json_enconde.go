package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Importante lembrar que para fazer a exportação de um struct, é necessário
// ter todas as keys com letra maiúscula.

// Qualquer coisa em go que esteja com letra minúscula não vai ser exportada.
// Exemplo func soma() -> não pode ser exportada por iniciar com letra minúscula.

type Pessoa struct {
	Nome          string  `json:"nome"`
	Sobrenome     string  `json:"sobrenome"`
	Idade         int     `json:"idade"`
	Profissao     string  `json:"profissao"`
	Contabancaria float64 `json:"contabancaria"`
}

// Nome string `json:"nome"` <- isso são tags. Tags são usadas para definir o nome da chave no json externamente.
// Dentro do programa em go, a struct Pessoa o nome da chave é Nome, porém, no json externo, o nome da chave é nome.
// Isso vale também para json recebidos externamente, quando o nome da chave vai ser sobrenome, por exemplo,
// mas internamente vai ser tratado como Sobrenome.

func main() {
	jamesbond := Pessoa{"James", "Bond", 35, "Agente de Segurança", 2340.0}

	enconder := json.NewEncoder(os.Stdout)
	enconder.Encode(jamesbond)

	// não vai funcionar porque o valor é printado direto no stdout do sistema e não salvo numa variável igual ao marshall
	enc := enconder.Encode(jamesbond)
	fmt.Println(enc)
}
