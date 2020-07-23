package main

import "fmt"

/*
	Constantes por definição não mudam, elas são as mesmas desde o inicio e é possivel iniciar
	constantes sem tipagem assim vc pode alocar nela o tipo quando for necessário.
*/

const (

	// Iotas guardam valores em contantes sem tipo definido e esses valores são sequenciais
	x = iota
	y = iota
	z = iota
)

func main() {
	fmt.Print(x, y, z)
	// ==> 0 1 2
}
