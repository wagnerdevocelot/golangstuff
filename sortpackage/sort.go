package main

import (
	"fmt"
	"sort"
)

/*
	Pacote sort já provide funções de ordenação prontas
*/

func main() {
	// slice de strings
	slaice := []string{"Wagner", "Beatriz", "Maria", "Valdemir", "Rodrigo"}
	sl := []int{3, 5, 6, 8, 9, 2334, 65, 9, 1, 0, -4}

	// metodo.função
	sort.Strings(slaice)
	fmt.Println(slaice)
	// saida ordenada por ordem alfabetica
	// ===> [Beatriz Maria Rodrigo Valdemir Wagner]

	sort.Ints(sl)
	fmt.Println(sl)
	// saida ordenada de numeros inteiros por ordem crescente
	// ===> [-4 0 1 3 5 6 8 9 9 65 2334]
}
