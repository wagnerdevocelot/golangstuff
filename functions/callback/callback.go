package main

import "fmt"

func main() {

	//y := numerosimpares(soma, []int{1, 2, 3, 4, 5}...)

	// usando o callback com uma função diferente que recebe parametros identicos
	z := numerosimpares(multiplica, []int{1, 2, 3, 4, 5}...)

	fmt.Println(z)
}

// função com parametro x variadico do tipo int e que retorna um int
func soma(x ...int) int {
	// n como agregador de valores
	n := 0
	// percorrendo os valores do argumento
	for _, v := range x {
		// somando os valores do argumento no agregador
		n += v
	}
	// retornando o agregador com a soma de todos os valores do argumento
	return n
}

// função praticamente igual a soma porém faz a multiplicação
func multiplica(w ...int) int {
	// aqui a variavel n tem o valor do indice 0 do slice do argumento
	// pois se fosse zero o unico resultado seria o proprio zero
	// escrevo isso obviamente depois de me interrogar por uns 4min porque a função não estava multiplicando kk
	n := w[0]
	// mesma coisa da função soma
	for _, v := range w {
		// mas aqui multiplica
		n *= v
	}
	return n
}

// uma função com dois argumentos (soma func(x ...int) int que por sí só tem seu proprio argumento e retorno)
// e y que é um inteiro que tambem retorna um int
// pode-se pensar que são 3 parametros, mas o x é um parametro da função soma e não da numerosimpares
func numerosimpares(soma func(x ...int) int, y ...int) int {
	// inicio de um alice de ints
	var slice []int
	// percorrendo o segundo parametro y
	for _, v := range y {
		// se o valor percorrido de y for impar (x%2 == 1 || x%2 != 0)
		if v%2 != 0 {
			// acrescente o valor de y na slice
			slice = append(slice, v)
		}
	}
	// total é o slice de numeros impares passado como argumento para a função no parametro
	// aqui usa-se soma somente para os fins do exemplo, o nome aqui é só um placeholder do callback
	// porem poderia ser qualquer outra função que recebe argumentos int variadicos e retorna int
	total := soma(slice...)
	// retorne o total de soma
	return total
}
