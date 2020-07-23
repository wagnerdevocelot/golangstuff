package main

import "fmt"

/*
	Operadores relacionais criam uma expressão que retornam um valor binario, true ou false
	Fazemos essas expressões utilizando operadores de relação que são:

	- Comparação: ==
	- Diferença: !=
	- Maior: >
	- Menor: <
	- Maior ou Igual: >=
	- Menor ou Igual: <=
*/

func main() {

	/* O necessario a se fazer aqui é fazer a expressão como pergunta (10 é igual a 9?)
	e ai retornar o valor verdadeiro ou falso*/

	a := (10 == 9) // false
	b := (10 != 9) // true
	c := (10 > 9)  // true
	d := (10 < 9)  // false
	e := (10 >= 9) // true
	f := (10 <= 9) // false

	fmt.Print(a, "\n", b, "\n", c, "\n", d, "\n", e, "\n", f)

}
