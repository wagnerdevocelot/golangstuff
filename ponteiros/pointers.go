package main

import "fmt"

// ponteiros são variáveis que contem um endereço da memória

func main() {
	// variaveis tem um valor guardado na memoria
	x := 10

	fmt.Println(&x) // através dessa notação é possível mostrar o espaço em memoria que (x) ocupa
	// ==> 0xc00001a0c0

	// aqui é o mesmo que dizer que (y) é o endereço de (x)
	y := &x

	// aqui (y) é um dereferencing que mostra o valor em memoria que existe em x
	fmt.Println(*y)
	// ===> 10

	// mostrando os tipos das variáveis
	fmt.Printf("%T, %T", x, y)
	// ===> int, *int

	// (x) é só um int, já (y) é um ponteiro que está apontando um espaço na memória que contem um int
}
