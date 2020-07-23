package main

import "fmt"

func main() {
	x := 10

	// uma função anonima não precisa de nome
	// esse x é igual ao x := 10 que ta ali em cima
	func(x int) {
		// aqui o numero 10 sendo multiplicado por 5
		fmt.Println(x * 5)
	}(x) // e esse tbm é o mesmo x, só que esse serve pra chamar a função
}

// funçoes anonimas são descartáveis
// nada demais para ver aqui
