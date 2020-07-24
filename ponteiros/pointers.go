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
	fmt.Printf("%T, %T\n\n", x, y)
	// ===> int, *int

	// (x) é só um int, já (y) é um ponteiro que está apontando um espaço na memória que contem um int

	/*
		O uso mais comum de ponteiros é para evitar o consumo de memoria passando valores para locais diferentes
		no programa, pelos valores em Go serem pass by value isso gera um custo alto em termos de performance.
	*/

	// duas variáveis com valores iguais
	z := 1
	w := 1

	passByvalue(z)   // passada como valor z++ = 2
	passBymemory(&w) // passada como ponteiro *w++ = 2

	fmt.Println(z) // porem na chamada da variável z ainda é = 1 é como se a função tivesse usado um dublê
	fmt.Println(w) // já w tem seu valor modificado diretamente na memoria
}

func passByvalue(w int) {
	w++
	fmt.Println("Variável passada como valor: ", w)

}

func passBymemory(w *int) {
	*w++
	fmt.Println("Variável passada como ponteiro: ", *w)
}
