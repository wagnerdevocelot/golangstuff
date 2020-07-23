package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

var wagner = pessoa{
	nome:  "wagner",
	idade: 29,
}

func main() {
	//fmt.Println(soma2(10, 4))
	//soma()
	//fmt.Println(retorna2(1, 2, 3, 4))
	//fmt.Println(retornaslice([]int{1, 2, 3, 4}...)) // usando unfrow pra desenrolar o slice no argumento
	//wagner.helloworld() // chamada de metodo
	//funcComoExpressao()

	/*
		// a função que retorna uma função precisa ser inserida em uma variavel antes pois nesse caso
		// a primeira função não tem argumento, e a segunda não tem nome
		y := retornaFunção()
		fmt.Println(y(10))

		// segunda forma de chamar uma função que retorna uma segunda função é usar dois parenteses
		fmt.Println(retornaFunção()(10))
	*/

}

/*
	Estrutura de uma função

	func (receiver) nome(parametro) (retorno) {o código em si}
*/

// função comum sem parametros
func soma() {
	// nesse caso 10 e 4 ficam hardcoded e a função em si é pouco dinamica
	fmt.Println(10 + 4)
}

// função comum com parametros
func soma2(x, y int) int {
	// nesse caso a função com dois parametros ja é muito mais dinamica pois podemos escolher os parâmtros
	// que mais nos interessa, assim se quisermos mudar os resultados mudados os parametros e não a função
	return x + y
}

// função que retorna dois valores inteiros atraves de um parametro váriadico
func retorna2(x ...int) (int, int) {
	soma := 0
	// iteração atraves dos valores passados como parametros adicionando os valores em sequencia
	for _, v := range x {
		soma += v
	}
	// retona o total da soma mais o tamaho do slice passado como parametro
	return soma, len(x)
}

func retornaslice(x ...int) int {
	soma := 0
	// iteração atraves dos valores passados como parametros adicionando os valores em sequencia
	for _, v := range x {
		soma += v
	}
	// retona o total da soma de um slice desenrolado
	return soma
}

// metodos tem um receiver que se refere ao tipo que irá utilizar a função criada e só esse tipo pode usá-la
// caso queira fazer operações modificarem atributos do receiver deve-se usar * junto ao struct (p *pessoa)
func (p pessoa) helloworld() {
	// a chamada do metodo funciona como chamada de metodos de objetos, nesse caso o parametro é o tipo
	// nome é o atributo do tipo
	fmt.Println(p.nome, "disse Olá mundo!")
}

func funcComoExpressao() {

	x := 10

	// colocando o tipo func dentro de uma variavel
	y := func(x int) int {

		return x * 5
	}

	fmt.Println(y(x))
}

// uma função que retorna uma nova função
func retornaFunção() func(int) int {
	// aqui dentro se declara func novamente, mas não se trata de 3 funções
	// aqui é um espelho da segunda função, se você tentar mudar o valor de int para string por exemplo
	// gera um erro dizendo que esperava um int
	return func(i int) int {
		return i * 10
	}
}
