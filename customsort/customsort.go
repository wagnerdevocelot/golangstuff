package main

import (
	"fmt"
	"sort"
)

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordernarPorPotencia []carro

func (x ordernarPorPotencia) Len() int {
	return len(x)
}

func (x ordernarPorPotencia) Less(i, j int) bool {
	return x[i].potencia < x[j].potencia
}

func (x ordernarPorPotencia) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type ordenarPorConsumo []carro
type ordenarPorLucro []carro

func main() {
	sonata := carro{
		nome:     "Sonata",
		potencia: 200,
		consumo:  2,
	}

	Golf := carro{
		nome:     "Golf",
		potencia: 300,
		consumo:  2,
	}

	Cobalt := carro{
		nome:     "Cobalt",
		potencia: 100,
		consumo:  2,
	}

	slaice := []carro{sonata, Golf, Cobalt}

	fmt.Println(slaice)
	sort.Sort(ordernarPorPotencia(slaice))
	fmt.Println(slaice)
}
