package main

import (
	// importação do pacote json referencia = https://golang.org/pkg/encoding/json/#pkg-examples
	"encoding/json"
	"fmt"
)

// structs que serão importados para json precisam ter os iniciais em letra maiuscula
type pessoa struct {
	Nome          string
	Sobrenome     string
	Idade         int
	Profissão     string
	ContaBancária float64
}

// objeto
var darthVader = pessoa{
	Nome:          "Darth",
	Sobrenome:     "Vader",
	Idade:         40,
	Profissão:     "Lord",
	ContaBancária: 109999,
}

func main() {
	// chamada e retorno da função em Json
	structoJSON(darthVader)
	// ==> {"Nome":"Darth","Sobrenome":"Vader","Idade":40,"Profissão":"Lord","ContaBancária":109999}
}

// função recebe struct pessoa como parametro
func structoJSON(p pessoa) {
	// função marshal devolve dois valores um erro e os dados em si
	// json.marchal recebe o parametro p da struct pessoa
	v, err := json.Marshal(p)
	// caso o erro seja diferente de nil printa o erro
	if err != nil {
		fmt.Println(err)
	}

	// saida da variável valor
	fmt.Println(string(v))
}
