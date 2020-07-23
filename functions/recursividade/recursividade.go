package main

import "fmt"

func main() {
	y := fatorial(3)
	z := loops(3)

	fmt.Println(y, z)
}

// a recursividade é o uso de uma função dentro dela mesma

// um calculo fatorial é como multiplicar do ultimo numero de uma lista até o primeiro
// 3 * 2 = 6
// 6 * 1 = 6
// logo o fatorial de 3 é 6, pois as multiplicações terminam no numero 1
func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}

// função recebe um inteiro devolve inteiro
// mesmo principio da função fatorial
func loops(x int) int {
	// contador iniciando do 1
	total := 1
	// enquanto x for maior que 1
	for x > 1 {
		// 1 x (x)
		total *= x
		// x = x - x
		x--
	}
	// retorna total
	return total
}
