package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

var wagner = pessoa{
	nome:  "Wagner",
	idade: 29,
}

func main() {
	/*
		- Crie um valor e atribua-o a uma variável.
		- Demonstre o endereço deste valor na memória.
	*/
	y := "Wagner"
	fmt.Println(&y)

	mudeMe(&wagner)
}

/*
- Crie um struct "pessoa"
- Crie uma função chamada mudeMe que tenha *pessoa como parâmetro.
- Essa função deve mudar um valor armazenado no endereço *pessoa.
    - Dica: a maneira "correta" para fazer dereference de um valor em um struct seria (*valor).campo
    - Mas consta uma exceção na documentação. Link: https://golang.org/ref/spec#Selectors
	- "As an exception, if the type of x is a named pointer type and (*x).f is a valid selector
	- expression denoting a field (but not a method),
    - → x.f is shorthand for (*x).f." ←
    - Ou seja, podemos usar tanto o atalho p1.nome quanto o tradicional (*p1).nome
- Crie um valor do tipo pessoa;
- Use a função mudeMe e demonstre o resultado.
*/

func mudeMe(p *pessoa) {
	p.idade++
	fmt.Println("Ano que vem", p.nome, "fará", p.idade, " anos")
}
