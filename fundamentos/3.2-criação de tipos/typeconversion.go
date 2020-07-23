package main

import "fmt"

// em Go é possivel criar novos tipos usando a keyword |type| nome do tipo | tipo do tipo

type qualquercoisa int

// Então você pode usar variaveis e outras estruturas baseadas no tipo criado

var numero qualquercoisa

func main() {
	// função que imprime o valor e tipo da variavel numero
	fmt.Printf("%v %T\n", numero, numero)
	// ==> 0 main.qualquercoisa

	// 0 (porque zero é o valor zero de um inteiro)
	// main.qualquercoisa (é o tipo criado)

	// em Go é possivel usar conversão ao contrario de java que usa coerção por exemplo
	// aqui iniciamos uma nova variavel que recebe o valor do nosso tipo criado convertido em float
	var numeroReal = float32(numero)

	fmt.Printf("%v %T", numeroReal, numeroReal)
	// ==> 0 float32

}
