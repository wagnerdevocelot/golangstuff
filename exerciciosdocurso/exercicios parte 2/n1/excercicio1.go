package main

import (
	"fmt"
)

func main() {
	hexaBinarioDecimal()
	operadores()
	constantes()
	deslocamentoDeBits()
	stringCrua()
	iotas()
}

func hexaBinarioDecimal() {
	// - Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
	x := 1337
	fmt.Printf("%d, %b, %#x", x, x, x)
	// https://play.golang.org/p/pVPw6cVXe77
}

func operadores() {
	/*
	   	- Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis.
	       - ==
	       - !=
	       - <=
	       - <
	       - >=
	       - >
	   	- Demonstre estes valores.
	*/

	var mimir1 int = 25
	var mimir2 int = 32

	a := (mimir1 == mimir2)
	b := (mimir1 != mimir2)
	c := (mimir1 <= mimir2)
	x := (mimir1 < mimir2)
	y := (mimir1 >= mimir2)
	z := (mimir1 > mimir2)

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", a, b, c, x, y, z)
	// https://play.golang.org/p/jbwOTjmp2Hw
}

func constantes() {
	/*
		- Crie constantes tipadas e não-tipadas.
		- Demonstre seus valores.
	*/
	const akira int = 2
	const tetsuo = 5
	fmt.Printf("%v,%T\n%v,%T", akira, tetsuo, akira, tetsuo)
	// https://play.golang.org/p/3H5ZVSo9l_o
}

func deslocamentoDeBits() {
	/*
			- Crie um programa que:
		    - Atribua um valor int a uma variável
		    - Demonstre este valor em decimal, binário e hexadecimal
		    - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
		    - Demonstre esta outra variável em decimal, binário e hexadecimal
	*/
	var numero int = 1337
	fmt.Printf("%d, %b, %#x\n", numero, numero, numero)
	alocamento := numero << 1
	fmt.Printf("%d, %b, %#x", alocamento, alocamento, alocamento)
	// https://play.golang.org/p/alXNDFRXXI1
}

func stringCrua() {
	/*
		- Crie uma variável de tipo string utilizando uma raw string literal.
		- Demonstre-a.
	*/
	var ascii string = `eae
	mano`

	fmt.Print(ascii)
	// https://play.golang.org/p/V3MYoYCN4Ch
}

func iotas() {
	/*
		- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
		- Demonstre estes valores.
	*/
	const (
		a = (iota + 2021)
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
	// https://play.golang.org/p/GloUGCA3Cdn
}
