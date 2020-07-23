package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//printa10k()
	//printUnicode()
	//loopVersary()
	//loopVersary2()
	//divisaopor4()
	//pikachu()
	//andreMarques()
	//suitche()
}

func printa10k() {
	// - Põe na tela: todos os números de 1 a 10000.
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
	// https://play.golang.org/p/o1bvbDyiZd1
}

func printUnicode() {
	// - Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
	for i := 65; i <= 90; i++ {
		fmt.Printf("%#U\n%#U\n%#U\n", i, i, i)

	}
	// https://play.golang.org/p/TxFfFEjuS3R
}

func loopVersary() {
	//- Crie um loop utilizando a sintaxe: for condition {}
	//- Utilize-o para demonstrar os anos desde que você nasceu.
	birthday := 1990
	for birthday <= 2020 {
		fmt.Println(birthday)
		birthday++
	}
	// https://play.golang.org/p/cKB7ExIkpqk
}

func loopVersary2() {
	// - Crie um loop utilizando a sintaxe: for {}
	// - Utilize-o para demonstrar os anos desde que você nasceu.
	birthday2 := 1990
	for {
		fmt.Println(birthday2)
		birthday2++
		if birthday2 > 2020 {
			break
		}
	}
	// https://play.golang.org/p/Zg5tF6xR3F9
}

func divisaopor4() {
	// - Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
	for i := 10; i <= 100; i++ {
		resto := i % 4
		fmt.Println(resto)
	}
	// https://play.golang.org/p/Q7BQAPO8ODf
}

func andreMarques() {
	// - Crie um programa que demonstre o funcionamento da declaração if.
	estiloNaoEMarra := true
	cabeça := estiloNaoEMarra
	if cabeça == true {
		fmt.Println("E ai marquin dj faz o sample de guitarra")
	}
	// https://play.golang.org/p/NfR0CbkNP5m
}

func pikachu() {
	// - Utilizando a solução anterior, adicione as opções else if e else.
	rand.Seed(time.Now().UnixNano())
	thunderPunch := rand.Intn(100)

	if thunderPunch > 80 {
		fmt.Println("Ataque Critico")
	} else if thunderPunch >= 30 {
		fmt.Println("Ataque Normal")
	} else {
		fmt.Println("Errou o ataque")
	}

	// https://play.golang.org/p/FbkBqiQktJl
}

func suitche() {
	// - Crie um programa que utilize a declaração switch,
	// onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
	esporteFavorito := "skate"
	switch esporteFavorito {
	case "skate":
		fmt.Println("Hey Ho, Lets Go!")
	default:
		fmt.Println("Quebrei a clavicula mano")
	}
	// https://play.golang.org/p/KKJaFYIYGJg
}
