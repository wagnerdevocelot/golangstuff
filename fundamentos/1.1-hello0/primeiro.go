// programas executaveis iniciam pelo pacote main
package main

/* Os coigos em go são organizados em packages
e para usá-los é necessario declarar a importação dos pacotes */
import "fmt"

/* a porta de entrada do programa é a função main blocos de código executaveis
serão feitos dentro dessa função */
func main() {
	/* O pacote fmt implementa I/O "input/output" formatada com funções análogas às printf e scanf de C.
	O formato'verbs' é derivado de C, mas é mais simples.*/

	fmt.Println("Hello ")
	fmt.Print("World! ")
}
