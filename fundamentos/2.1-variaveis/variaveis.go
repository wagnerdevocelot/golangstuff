package main

import "fmt"

func main() {

	// https://www.portaldoholanda.com.br/amizade/menino-faz-surpreendente-amizade-com-marmota

	/* o Go é uma linguagem de tipagem estatica e geralmente se precisa de declaração do tipo nesse
	caso a variavel x é um inteiro e por definição ele vale 0 pois não lhe foi atribuido nenhum valor */
	var x int
	fmt.Println(x)

	// esse é o perador gopher ou marmotinha se preferir
	/* com a marmotinha vc poder inicializar a variavel y ja declarando que seu valor é 10 é uma declaração
	mais dinâmica e por isso a marmotinha bilha no go */
	y := 10
	fmt.Println(y)

	/* como o go é fortemente tipado ainda assim redeclarar uma variavel inicializada como inteiro passando
	uma string ira retornar um erro, é interessante pois permanece a tipagem forte mas você pode inicializar
	dinamicamente */
	//y = "10"
	fmt.Println(y)

	// ==> cannot use "10" as type int in assigment

	/* mesmo ja tendo inicializado a variavel com a marmotinha vc pode reatribuir com o '=' que é o operador
	de atribuição, isso desde que o novo valor seja do mesmo valor da primeira inicialização */
	y = 20
	fmt.Println(y)

	/* a marmotinha ela só funciona no contexto de inicialização, por isso quando queremos reatribuir um novo
	valor usamos '=' mas se vc quer reatribuir e ja inicializar uma nova variável (nesse caso o 'z') funciona */
	y, z := 10, "gopher"
	fmt.Println(y, z)

	/* a marmotinha ela só pode ser usada dentro de um bloco de código, variavéis definidas de forma comum como
	var x int podem ser definidas fora da func main() o que dará a elas uma abrangencia maior a nivel de package

	geralmente em um projeto usamos mais de um arquivo go para manter uma certa organização de código no projeto
	variaveis definidas fora do func main() poderão ser usadas inclusive fora desse mesmo arquivo

	a marmotinha é usada ate onde func main() abrange apos a sua inicialização, como não podemos ter variaveis
	sem proposito devemos ter uma função que as utilize e essa função vem sempre depois da variavel pois
	não podemos usá-la em um contexto onde ela nem existia antes
	o compilador go vai ler o seu código de cima pra baixo até chegar na ultima linha se a fnção vier antes da
	variavel ele vai te responder com erro*/

	/* variaveis definidas globalmente não precisam ter valor atribuido, vc pode ter uma (var numero int)
	que não recebe nada, mas ela precisa ser atribuida dentro de um bloco de código e somente dentro dentro
	pra quem ja vu C ou pascal sabe como funciona */
}

// quando se define a variavel fora de um bloco de código o go não te obriga a usar ela ^^
var x = 4
