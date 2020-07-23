package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

var wagner = pessoa{
	nome:      "Wagner",
	sobrenome: "Abrantes",
	idade:     29,
}

func main() {
	//fmt.Println(chamaOInteiro())
	//fmt.Println(retornaStringEInt())
	//fmt.Println(exe2([]int{1, 2, 3, 4, 5}...))
	//fmt.Println(soma([]int{1, 2, 3, 4, 5, 6, 6, 6}))
	//novosMutantes()
	//wagner.personalIdentifier()
	//info(dominó)
	//info(bolacha)
	//backToTheFuture()(1977)
	//a := sliceMaiusculo(maiuscula, []string{"a", "b", "c", "d"}...)
	//fmt.Println(a)

	//quote()()

	/*
			- Crie e utilize uma função anônima.

		y := []int{1, 2, 3, 4, 5}

		func(y []int) int {
			total := 0
			for _, v := range y {
				total += v
			}
			fmt.Println(total)
			return total
		}(y)
	*/

	//x(6)

}

/*
- Exercício:
    - Crie uma função que retorne um int
    - Crie outra função que retorne um int e uma string
    - Chame as duas funções
    - Demonstre seus resultados
*/

func chamaOInteiro() int {
	x := 10
	y := 5
	return x / y
}

func retornaStringEInt() (string, int) {
	x := 10
	y := 5
	return "A divisão é: ", x / y
}

/*
- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e
- retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
*/

func exe2(x ...int) int {
	soma := 0
	for _, v := range x {
		soma += v
	}
	return soma
}

func soma(x []int) int {
	contador := 0
	for _, v := range x {
		contador += v
	}
	return contador
}

/*
- Utilize a declaração defer
- de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
*/

func novosMutantes() {
	defer fmt.Println("Laçamento Hoje")
	fmt.Println("O filme foi adiado novamente")
}

/*
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
*/

func (p pessoa) personalIdentifier() {
	fmt.Println("Olá,", p.nome, p.sobrenome, "Seu acesso está garantido, parabens pelo seu anivesário de", p.idade, "anos")
}

/*
- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
*/
type quadrado struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}
type figura interface {
	area() float64
}

var dominó = quadrado{4, 8}
var bolacha = circulo{10}

func (q quadrado) area() float64 {
	return q.altura * q.largura
}

func (c circulo) area() float64 {
	return 2 * 3.14 * c.raio
}

func info(f figura) {
	fmt.Println(f.area())
}

/*
	- Atribua uma função a uma variável.
	- Chame essa função.
*/

var x = func(x int) {
	if x%2 == 0 {
		fmt.Printf("%v, é um número par", x)
	} else {
		fmt.Printf("%v, é um número par", x)
	}
}

/*
- Crie uma função que retorna uma função.
- Atribua a função retornada a uma variável.
- Chame a função retornada.
*/

func backToTheFuture() func(int) {
	return func(y int) {
		fmt.Printf("We are in %v holly shit", y)
	}
}

/*
- Callback: passe uma função como argumento a outra função.
*/

func maiuscula(s string) string {
	grande := strings.ToUpper(s)
	return grande
}

func sliceMaiusculo(ma func(x string) string, s ...string) []string {
	var slice []string
	for _, v := range s {
		maiuscula := ma(v)
		slice = append(slice, maiuscula)
	}
	return slice
}

/*
- Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função,
- onde esta outra função faz uso de uma variável alem de seu scope.
*/

func quote() func() {
	characther := "Neo"

	return func() {
		fmt.Printf("I don’t like the idea that I’m not in control of my life. - %v", characther)
	}
}
