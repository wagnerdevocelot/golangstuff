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

}
