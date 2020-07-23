package main

// pacotes importados automaticamente
import (
	"fmt"
	"math"
)

func main() {
	// definindo uma constante passando o tipo
	const PI float64 = 3.1415
	// definindo a variavel com o tipo inferido
	var raio = 3.2

	// definição e atribuição ao mesmo tempo
	area := PI * math.Pow(raio, 2)
	// chamando a variavel criada acima, pois uma variavel que não é usada gera um erro de compilação
	fmt.Println("A área da circunferência é:", area)

	// definição de variaveis e constantes atraves de blocos
	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// definindo mais de uma variavel em uma mesma linha
	var e, f bool = true, false
	fmt.Println(e, f)

	// definindo vriaveis de tipos diferentes na mesma linha usando operador de atribuição
	g, h, i := 2, false, "opa!"
	fmt.Println(g, h, i)
}
