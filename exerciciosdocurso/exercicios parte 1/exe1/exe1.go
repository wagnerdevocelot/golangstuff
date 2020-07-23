package main

import (
	"fmt"
)

func main() {
	marmotinha()
	scopo()
	typeCreate()
	passaStrings()
	conversão()
}

func marmotinha() {
	/*
		- Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".
		    1. 42
		    2. "James Bond"
		    3. true
		- Agora demonstre os valores nestas variáveis, com:
		    1. Uma única declaração print
		    2. Múltiplas declarações print */
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	// https://play.golang.org/p/EGj-XF3lj48
}

var x int
var y string
var z bool

func scopo() {
	/*
		- Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:
		    1. Identificador "x" deverá ter tipo int
		    2. Identificador "y" deverá ter tipo string
		    3. Identificador "z" deverá ter tipo bool
		- Na função main:
		    1. Demonstre os valores de cada identificador
		    2. O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
		- Solução: https://play.golang.org/p/pAFd-F7uGZ
	*/

	fmt.Print(x, "\n", y, "\n", z)

	/* Quando variáveis são definidas e não inicializadas elas tem um valor padrão que é o valor zero então
	   cada tipo possui seu valor default */

	// https://play.golang.org/p/99L1zwZhrr3
}

var j = 42
var k = "James Bond"
var l = true

func passaStrings() {
	/*
		- Utilizando a solução do exercício anterior:
		    1. Em package-level scope, atribua os seguintes valores às variáveis:
		        1. para "x" atribua 42
		        2. para "y" atribua "James Bond"
		        3. para "z" atribua true
		    2. Na função main:
		        1. Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
		        2. Demonstre a variável "s".
	*/
	s := fmt.Sprintln(j, k, l)
	fmt.Println(s)
	// https://play.golang.org/p/408_ZnVaMu0
}

func typeCreate() {
	/*
		- Crie um tipo. O tipo subjacente deve ser int.
		- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
		- Na função main:
		    1. Demonstre o valor da variável "x"
		    2. Demonstre o tipo da variável "x"
		    3. Atribua 42 à variável "x" utilizando o operador "="
		    4. Demonstre o valor da variável "x"
		- Para os aventureiros: https://golang.org/ref/spec#Types
		- Agora já somos quase ninjas nível 1!
	*/
	type opa int

	var erros opa

	fmt.Printf("%v, %T \n", erros, erros)
	erros = 42
	fmt.Println(erros)
	// https://play.golang.org/p/MUH8NRoCa0d
}

func conversão() {
	/*
		- Utilizando a solução do exercício anterior:
		    1. Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
		    2. Na função main:
		        1. Isto já deve estar feito:
		            1. Demonstre o valor da variável "x"
		            2. Demonstre o tipo da variável "x"
		            3. Atribua 42 à variável "x" utilizando o operador "="
		            4. Demonstre o valor da variável "x"
		        2. Agora faça tambem:
		            1. Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
		            2. Demonstre o valor de "y"
		            3. Demonstre o tipo de "y"
	*/
	type eita int

	var a eita
	var b int

	fmt.Printf("%v, %T \n", a, a)
	a = 42
	fmt.Println(a)
	b = int(a)
	fmt.Printf("%v, %T \n", b, b)
	// https://play.golang.org/p/NC__5ElFVdd
}
