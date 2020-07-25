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

	// metodo.função
	sort.Strings(slaice)
	fmt.Println(slaice)
	// saida ordenada por ordem alfabetica
	// ===> [Beatriz Maria Rodrigo Valdemir Wagner]
}
