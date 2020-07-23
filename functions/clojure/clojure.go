package main

import "fmt"

func main() {

	a := i()

	// nesse segundo caso a := i() cria um escopo novo e as funções rodam dentro  desse contexto
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	// chamando a função sem uma variável ela funciona como um clojure em cada chamada
	// aqui cada chamada cria um contexto novo
	fmt.Println(i()())
	fmt.Println(i()())
	fmt.Println(i()())
}

// função que retorna função
// nenhuma das duas espera argumento
// um retorno int
func i() func() int {
	// contador x
	x := 0
	// função como retorno
	return func() int {
		// adição ao contador x
		x++
		// retorno do valor
		return x
	}
}
