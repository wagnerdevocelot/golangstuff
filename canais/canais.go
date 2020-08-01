// - Canais são o Jeito Certo® de fazer sincronização e código concorrente.
// - Eles nos permitem trasmitir valores entre goroutines.
// - Servem pra coordenar, sincronizar, orquestrar, e buffering.

package main

import "fmt"

func main() {

	// exemplo não funciona pois não tem uma go routine fazendo cominicação bidirecional com o canal
	canal := make(chan int)
	//canal <- 42
	//fmt.Println(canal)

	go func() {
		canal <- 42
	}()

	fmt.Println(<-canal)
	// ==> 42

	// um canal pode ser definido com um buffer para que a go func não precise levar todos os dados
	// porem aqui só um valor pode ser ignorado, mais de um resulta em deadlock e buffers em geral
	// não são encorajados em golang
	canal2 := make(chan int, 1)
	canal2 <- 50
	fmt.Println(<-canal2)
	// ==> 50

	// ========================== canais direcionais ================================

	// é possivel ter canais que recebem e canais que só enviam dados

	canaldirecional := make(chan int)

	// pelo menos uma das funções deve ser uma go func para que as funções rodem em paralelo

	go send(canaldirecional)

	receiver(canaldirecional)
	// ==> 33

	// conversão de canais pode ser feita usando a notação de conversão de tipos
	// Porem isso só pode ser feito a partir de canais bidirecionais
	conversãoSend := (<-chan int)(canaldirecional)
	fmt.Printf("%T\n", conversãoSend)
	// ==> <-chan int
	conversãoReceiver := (chan<- int)(canaldirecional)
	fmt.Printf("%T\n", conversãoReceiver)
	// ==> chan<- int

	// ========================== range & close ================================

	forCanal := make(chan int)

	go loopSend(10, forCanal)

	// é possivel usar range sobre um canal, porém não tem (i) apenas value(v)
	for v := range forCanal {
		fmt.Println("recebido do canal:", v)
	}
	// recebido do canal: 0
	// recebido do canal: 1
	// recebido do canal: 2
	// recebido do canal: 3
	// recebido do canal: 4
	// recebido do canal: 5
	// recebido do canal: 6
	// recebido do canal: 7
	// recebido do canal: 8
	// recebido do canal: 9
	// recebido do canal: 10

}

// send usa um canal bidirecional e transforma ele em um canal que envia dados
func send(s chan<- int) {
	s <- 33
}

// receiver usa um canal bidirecional como parametro e o transforma em um receptor de dados
func receiver(r <-chan int) {
	fmt.Println(<-r)
}

// loopSend faz envios de (i) para o chanel(s) recebido em parametro, iterando sobre o valor de (t)
func loopSend(t int, s chan<- int) {
	for i := 0; i <= t; i++ {
		s <- i
	}
	// necessário fechar um canal após a iteração, caso não seja fechado um range pode ficar esperando
	// um novo valor quando não existem mais valores e dar erro de deadlock
	close(s)
}
