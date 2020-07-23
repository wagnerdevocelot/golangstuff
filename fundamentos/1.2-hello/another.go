// programas executaveis iniciam pelo pacote main
package main

/* Os coigos em go são organizados em packages
e para usá-los é necessario declarar a importação dos pacotes */
import "fmt"

/* a porta de entrada do programa é a função main blocos de código executaveis
serão feitos dentro dessa função */
func main() {

	/* A função println além de retornar os parametros enviados ela implementa atraves da interface
	o numero de bytes escritos e os erros, caso não haja erro o retorno será nill */

	numerodebytes, erros := fmt.Println("Hello world", "Ola mundo", 100)
	fmt.Println(erros)
	fmt.Println(numerodebytes)

	/* Go não permite variáveis sem proposito, então caso ela não esteja sendo utilizada ele retorna um erro */

	n, e := fmt.Println("Hello world", "Ola mundo", 100)
	fmt.Println(e)

	// ==> another.go:21:2: n declared and not used

	/* quando usamos o underscore no lugar onde deveria vir a primeira variável que nesse caso é
	a de numero de bytes estamos dizendo que queremos ignorá-la e ai só é impresso o segundo valor
	que nesse exemplo é o errors que é nill */
	_, f := fmt.Println("Hello world", "Ola mundo", 100)
	fmt.Println(f)

}
