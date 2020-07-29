/*
	Goroutines são a forma de executar funções usando o máximo do processador, se o processador possuir
	1 core, provavelmente as funções irão executar em concorrencia, e com mais cores em paralelo
	Ressaltar que concoreencia e paralelismo não são a mesma coisa
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// como as funções perdem a corrida contra a função main na execução é preciso o uso de metodos para que
// a execução não termine antes das demais goroutines acabarem de executar
var wg sync.WaitGroup

func main() {

	// Add precisa dizer quantas goroutines serão executadas
	wg.Add(2)

	// adicionando go antes da função ela ja se torna concorrente
	// porem a propria função main é uma função concorrente, então se a função main terminar antes
	// essas duas funções não são executadas
	go função1()
	go função2()

	// Essa é a saida das funções, repare que não é uniforme, cada função executou uma pequena parte
	// como se tivessem disputando memoria
	/*
		Func2: 0
		Func2: 1
		Func2: 2
		Func2: 3
		Func2: 4
		Func1: 0
		Func1: 1
		Func1: 2
		Func1: 3
		Func1: 4
		Func1: 5
		Func1: 6
		Func1: 7
		Func1: 8
		Func1: 9
		Func2: 5
		Func2: 6
		Func2: 7
		Func2: 8
		Func2: 9
	*/

	// runtime.NumCPU imprime no stdout o numero de cores do processador
	fmt.Println(runtime.NumCPU())
	// ==> 4 no meu caso

	// runtime.NumGoroutine imptime no stdout o numero de goroutnes sendo executadas
	fmt.Println(runtime.NumGoroutine())
	// ==> 3 goroutines sendo executadas nesse programa /main/função1/função2

	// wait diz para que a função main espere a execução das goroutines
	wg.Wait()
}

func função1() {
	for i := 0; i < 10; i++ {
		fmt.Println("Func1:", i)
	}
	// Done avisa main quando a função termina de ser executada
	wg.Done()
}

func função2() {
	for i := 0; i < 10; i++ {
		fmt.Println("Func2:", i)
	}
	// Done avisa main quando a função termina de ser executada
	wg.Done()
}
