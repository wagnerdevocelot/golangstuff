package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//idade()
	//aluno()
	//quadrado()
	//cond()
	//count()
	//reloginho()
	//exe1()
	//exe2()
	//exe3()
	//exe4()
	//exe5()
	//exe6()
	//soma()
	//Version()
	//produto()
	//media()
}

func idade() {
	// mostra idade e nome
	var idade int
	var nome string
	fmt.Println("Seu nome: ")
	fmt.Scanf("%v", &nome)
	fmt.Println("Digite a sua idade: ")
	fmt.Scanf("%v", &idade)
	fmt.Printf("%v, sua idade é %v anos \n", nome, idade)
}

func aluno() {
	// mostra o nome e a media de nota de um aluno
	var nome string
	var nota1, nota2, nota3 float64
	fmt.Println("Qual o nome do aluno?: ")
	fmt.Scanf("%v", &nome)
	fmt.Println("Quais as notas desse semestre?: ")
	fmt.Scanf("%v\n%v\n%v", &nota1, &nota2, &nota3)
	media := (nota1 + nota2 + nota3) / 3
	fmt.Printf("O %v tem a média de %2.f nesse semestre", nome, media)
}

func quadrado() {
	// um numero elevado ao quadrado
	var numero int
	fmt.Println("Digite um numero: ")
	fmt.Scanf("%v", &numero)
	fmt.Printf("O %v elevado ao quadrado resulta em %v\n", numero, numero*numero)
}

func cond() {
	// condicional simples
	var inteiro int
	fmt.Println("Digite sua idade: ")
	fmt.Scanf("%v", &inteiro)
	if inteiro > 70 {
		fmt.Println("Ta vei em fi")
	} else if inteiro > 21 {
		fmt.Println("Daqui pra frente tudo piora")
	} else {
		fmt.Println("Leite com pêra")
	}
}

func count() {
	// for simples
	for i := 1; i <= 100; i++ {
		fmt.Printf("%v\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func reloginho() {
	// for aninhado para iterar em horas e minutos
	for h := 0; h <= 12; h++ {
		fmt.Println(h, "Horas")
		for m := 0; m <= 60; m++ {
			fmt.Print(" ", m)
			fmt.Print("Minutos")
		}
		fmt.Print("\n")
	}
}

func exe1() {
	// 1. Write a C program to print your name, date of birth. and mobile number.
	fmt.Println("Name   : Alexandra Abramov ")
	fmt.Println("DOB    : July 14, 1975 ")
	fmt.Println("Mobile : 99-9999999999")
}

func exe2() {
	// Write a C program to print a block F using hash (#), where the F has a height of six characters
	//and width of five and four characters. And also to print a big 'C'
	fmt.Println(`
		######
		#
		#
		#####
		#
		#
		#
		
		  ######
		##      ##
		#
		#
		#
		#
		#
		##      ##
		  ######`)
}

func exe3() {
	// Write a C program to print the following characters in a reverse way
	a := "X"
	b := "M"
	c := "L"

	fmt.Printf("The reverse of %v%v%v is %v%v%v\n", a, b, c, c, b, a)
}

func exe4() {
	// Write a C program to compute the perimeter and area
	// of a rectangle with a height of 7 inches. and width of 5 inches
	altura := 7
	largura := 5

	perimetro := 2 * (altura + largura)
	fmt.Printf("O perimetro do retangulo é de: %v metros\n", perimetro)

	area := altura * largura
	fmt.Printf("A area do retangulo é de: %v metros\n", area)
}

func exe5() {
	// Write a C program to compute the perimeter and area of a circle with a radius of 6 inches.
	var area, perimetro float64
	const pi float64 = 3.14
	raio := 6
	raiof := float64(raio)

	perimetro = (2 * pi) * raiof
	fmt.Printf("O perimetro do circulo é: %2.f metros\n", perimetro)

	area = pi * raiof * raiof
	fmt.Printf("A area do circulo é: %2.f metros\n", area)

}

func exe6() {
	// Write a C program to convert specified days into years, weeks and days.
	dias := 1329

	anos := dias / 365
	semanas := (dias % 365) / 7
	dias = dias - ((anos * 365) + (semanas * 7))
	fmt.Println(anos)
	fmt.Println(semanas)
	fmt.Println(dias)

}

func soma() {
	// Write a C program that accepts two integers from the user and calculate the sum of the two integers
	var x, y int
	fmt.Println("Digite dois numeros que deseja somar: ")
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)
	fmt.Printf("A soma de %v e %v é igual a: %v\n", x, y, x+y)
}

func Version() string {
	// retorna a versão atual do go
	return runtime.Version()
}

func produto() {
	// Write a C program that accepts two integers from the user and calculate the product of the two integers.
	var x, y int
	fmt.Println("Digite dois numeros que deseja multiplicar: ")
	fmt.Scanf("%d", &x)
	fmt.Scanf("%d", &y)
	fmt.Printf("A soma de %v e %v é igual a: %v\n", x, y, x*y)
}

func media() {
	var peso1, qtd1, peso2, qtd2, resultado float64

	fmt.Println("Digite o peso do primeiro item: ")
	fmt.Scanf("%v", &peso1)
	fmt.Println("Digite a quantidade do primeiro item: ")
	fmt.Scanf("%v", &qtd1)
	fmt.Println("Digite o peso do segundo item: ")
	fmt.Scanf("%v", &peso2)
	fmt.Println("Digite a quantidade do segundo item: ")
	fmt.Scanf("%v", &qtd2)
	resultado = ((peso1 * qtd1) + (peso2 * qtd2)) / (qtd1 + qtd2)
	fmt.Printf("A média dos valores é de: %v\n", resultado)

}
