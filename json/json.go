package main

import (
	"encoding/json"
	"fmt"
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
	darthvader := Pessoa{"Luki", "Skywalker", 57, "Jedi", 7340.0}

	j, err := json.Marshal(jamesbond)

	if err != nil {
		fmt.Println(err)
	}

	e, err := json.Marshal(darthvader)

	darthvaderConvertido := string(e)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(j))
	fmt.Println(darthvaderConvertido)

	// O trecho acima converte uma struct para json usando o marshal.

	// O trecho abaixo converte um json para uma struct usando o unmarshal.

	// Para converter o conteúdo de uma variável para slice de bytes, é necessário
	// usar o Sprintf, uma vez que ele vai uma string como retorno
	sb := []byte(fmt.Sprintf(`%v`, darthvaderConvertido))

	var jbond Pessoa // crio uma variável do tipo Pessoa

	// passo o slice de bytes para a variável jbond. O retorno disso é um erro.
	// Caso não haja erro, o resultado é jogado para dentro da variável referenciada em memória
	err = json.Unmarshal(sb, &jbond)

	if err != nil {
		fmt.Println(err)
	}

	// A saída vai ser o conteúdo da struct informada
	fmt.Println(jbond, jbond.Nome)

}
