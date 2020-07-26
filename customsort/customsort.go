package main

import (
	"fmt"
	"sort"
)

// struct de carro
type carro struct {
	nome     string
	potencia int
	consumo  int
}

// tipo ordenar por (x) slice do struct carro
type ordernarPorPotencia []carro

// metodo receiver do tipo ordenarporpotencia, retorna um int
func (x ordernarPorPotencia) Len() int {
	// aqui ele define no metodo de sort qual o tamanho do slice que vai ser ordenado
	return len(x)
}

// metodo less retorna um boolean pra saber se o indice i é maior que o indice j
func (x ordernarPorPotencia) Less(i, j int) bool {
	return x[i].potencia > x[j].potencia
}

// Swap faz a inversão dos valores de i e j caso a condição de Less seja verdadeira
func (x ordernarPorPotencia) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type ordenarPorConsumo []carro

// a mesma implementação com um atributo diferente
// nesse caso ordenado pelos casso com maior consumo
func (x ordenarPorConsumo) Len() int           { return len(x) }
func (x ordenarPorConsumo) Less(i, j int) bool { return x[i].consumo > x[j].consumo }
func (x ordenarPorConsumo) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type ordenarPorEconomia []carro

// invertendo a lógica temos a ordenação por economia
func (x ordenarPorEconomia) Len() int           { return len(x) }
func (x ordenarPorEconomia) Less(i, j int) bool { return x[i].consumo < x[j].consumo }
func (x ordenarPorEconomia) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {
	// objetos de carros
	sonata := carro{
		nome:     "Sonata",
		potencia: 200,
		consumo:  7,
	}

	Golf := carro{
		nome:     "Golf",
		potencia: 300,
		consumo:  9,
	}

	Cobalt := carro{
		nome:     "Cobalt",
		potencia: 100,
		consumo:  2,
	}

	// slice com os onjetos
	slaice := []carro{sonata, Golf, Cobalt}

	fmt.Println("Valores originais: ", slaice)
	sort.Sort(ordernarPorPotencia(slaice))
	fmt.Println("Valores ordenados por potencia do maior para o menor: ", slaice)
	sort.Sort(ordenarPorConsumo(slaice))
	fmt.Println("Valores ordenados por consumo do maior para o menor: ", slaice)
	sort.Sort(ordenarPorEconomia(slaice))
	fmt.Println("Valores ordenados por economia do menor para o maior: ", slaice)
}
