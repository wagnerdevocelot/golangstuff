package main

import "fmt"

// Structs são agrupamentos de dados de tipos diferentes
type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	//struct1()
	//structaninhado()
	//anonnymousStructs()
}

func struct1() {
	// delcaração explicita
	cliente1 := cliente{
		nome:      "Wagner",
		sobrenome: "Abrantes",
		fumante:   true,
	}

	// delcaração mais concisa, sempre usar os dados na ordem correta
	// necessario verificar a declaração global para leitura
	cliente2 := cliente{"Wagner", "Abrantes", true}

	fmt.Println(cliente1, "\n", cliente2)
}

func structaninhado() {
	/*
		O Struct aninhado leva dentro dele um ou mais structs, lembra bastante a composição de dados
		da Orientação a Objetos com objetos que tem atributos diferentes recebidos de uma herança
	*/

	pessoa1 := profissional{
		pessoa: pessoa{
			nome:  "Wagner",
			idade: 29,
		},
		titulo:  "Developer",
		salario: 6000,
	}

	// declaração implicita
	pessoa2 := profissional{pessoa{"Wagner", 29}, "Developer", 6000}

	// a forma de acesso dos atributos do struct funcionam da mesma forma que em objetos
	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.nome)
}

func anonnymousStructs() {
	// não fica muito claro, mas é perceptivel que o uso é mais simples
	x := struct {
		nome  string
		idade int
	}{
		nome:  "Wagner",
		idade: 29,
	}

	fmt.Println(x)
}
