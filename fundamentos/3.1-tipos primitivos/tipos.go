package main

import "fmt"

func main() {
	// aqui uso o printf para formatar a saida esse (%v, %T) é usado para imprimir tanto a variavel como o valor
	fmt.Printf("%v, %T", numeroInteiro, numeroInteiro)
	fmt.Printf("%v, %T", numeroReal, numeroReal)
	fmt.Printf("%v, %T", conjuntoDeCaracteres, conjuntoDeCaracteres)
	fmt.Printf("%v, %T", booleano, booleano)
	fmt.Printf("%v, %T", inteiro, inteiro)

	/* Quando variáveis são definidas e não inicializadas elas tem um valor padrão que é o valor zero então
	   cada tipo possui seu valor default */

}

// tipos primitivos

var numeroInteiro int

// um numero inteiro que pode ser ou não negativo
// valor zero = 0

var numeroReal float64

// um numero com casas decimais
// valor zero = 0.0

var conjuntoDeCaracteres string

/* Strings são definidas por aspas duplas em geral e dentro delas vão desde letras a numeros e caracteres
   especiais como acentuação, pontos, underscore, asteristico e etc você pode usar números como strings
   por exemplo quando a varável recebe um numero de telefone ou um cpf porque você não irá fazer
   cálculos com eles */

// valor zero = "" se trata de uma string só que vazia

var booleano bool

/* O tipo boolean é um sistema de lógica algébrica definido por George Boole no século XIX ele retorna
(true) ou (false) como resultado de uma expressão */

// valor zero = false

/* Geralmente no valor booleano o false é o defalut até que você mude ele para true tornando uma condição
   verdadeira. Um exemplo muito usado:
   Um site pode ser acessado por um administrador ou um user comum que teria acesso a somente algumas
   informações enquanto o administrador teria todas ou as que fossem pertinentes

   #Mostrar
   A
   B
   C

   se admin for true?

   #Mostrar
   A
   C
   D
   E

*/

/*
   Em Go é possivel definir na criação das variaveis além dos tipos, o tamanho de espaço que aquele tipo
   armazena, pode ser 8 bits, 16, 32, 64

   então dependendo de como fazemos uso dessas variáveis podemos encontrar dificuldades como Overflow

   um uint16 vai de 0 a 65535, então se eu tenho uma variavel que recebe + 1 toda vez que eu executo um
   código quer dizer que quando ela chegar nesse máximo (65535) ela não vai conseguir mais computar um
   novo acrescimo e vai imprimir zero novamente.

   OBS. Isso foi o que causou o Bug do milénio no ano 2000
*/

var inteiro uint64
