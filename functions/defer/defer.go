package main

import "fmt"

func main() {
	// defer é um statement que prorroga uma funcionalidade ou um bloco de código

	defer fmt.Println("Wagner") // esse aqui apesar de vir primeiro será executado depois
	fmt.Print("Meu nome é: ")   // e esse aqui vem primeiro

} // o defer vai ser executado quando a leitura do código chegar aqui no final do code block

/*
	Nos exemplos citados ele pode ser usado como uma forma de fechar arquivos antes de dar overload
	fechar conexões web dentre outras opções
*/
