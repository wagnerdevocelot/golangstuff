package main

import (
	"fmt"
)

func main() {
	//arrei()
	//slaice()
	//fatiaDeSlice()
	//appendar()
	//slippend()
	//makingmake()
	//slice2d()
	mapa()
}

func arrei() {
	/*
		- Usando uma literal composta:
		- Crie um array que suporte 5 valores to tipo int
		- Atribua valores aos seus índices
		- Utilize range e demonstre os valores do array.
		- Utilizando format printing, demonstre o tipo do array.
	*/
	arrei := [5]int{10, 20, 30, 40, 50}

	for _, value := range arrei {
		fmt.Printf("%v %T\n", value, value)
	}

	// https://play.golang.org/p/88wlbZnq4JA
}

func slaice() {
	/*
		- Usando uma literal composta:
		- Crie uma slice de tipo int
		- Atribua 10 valores a ela
		- Utilize range para demonstrar todos estes valores.
		- E utilize format printing para demonstrar seu tipo.
	*/

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Printf("%T", slice)

	// https://play.golang.org/p/ZImmPt9BUrR
}

func fatiaDeSlice() {
	/*
		- Utilizando como base o exercício anterior, utilize slicing para demonstrar os valores:
		- Do primeiro ao terceiro item do slice (incluindo o terceiro item!)
		- Do quinto ao último item do slice (incluindo o último item!)
		- Do segundo ao sétimo item do slice (incluindo o sétimo item!)
		- Do terceiro ao penúltimo item do slice (incluindo o penúltimo item!)
		- Desafio: obtenha o mesmo resultado acima utilizando a função len() para determinar o penúltimo item
	*/

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println(slice[:3])
	fmt.Println(slice[4:])
	fmt.Println(slice[1:7])
	fmt.Println(slice[2:9])
	fmt.Println(slice[2 : len(slice)-1])

	// https://play.golang.org/p/tPRbMX4jo6h
}

func appendar() {
	/*
		- Começando com a seguinte slice:
		- x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		- Anexe a ela o valor 52;
		- Anexe a ela os valores 53, 54 e 55 utilizando uma única declaração;
		- Demonstre a slice;
		- Anexe a ela a seguinte slice:
		- y := []int{56, 57, 58, 59, 60}
		- Demonstre a slice x.
	*/

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)
	x = append(x, 53, 54, 55)

	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}

	x = append(x, y...)

	fmt.Println(x)
	// https://play.golang.org/p/p1nbW2Cfe7P
}

func slippend() {
	/*
				- Comece com essa slice:
		    - x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
		- Utilizando slicing e append, crie uma slice y que contenha os valores:
		    - [42, 43, 44, 48, 49, 50, 51]
	*/

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(x[:3], x[6:len(x)]...)

	fmt.Println(y)
	// https://play.golang.org/p/D7Nrc6kotAE
}

func makingmake() {
	/*
				- Crie uma slice usando make que possa conter todos os estados do Brasil.
			- Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás",
			 "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná",
			  "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia",
			   "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
		- Demonstre o len e cap da slice.
		- Demonstre todos os valores da slice *sem utilizar range.*
	*/
	y := make([]string, 26, 26)

	y = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás",
		"Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná",
		"Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia",
		"Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(y), cap(y))
	fmt.Println(y)

	for i := 0; i < len(y); i++ {
		fmt.Println(y[i])
	}
	// https://play.golang.org/p/4lzXEHgv4Fg
}

func slice2d() {
	/*
				- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice
				multi-dimensional da seguinte maneira:
		    - "Nome", "Sobrenome", "Hobby favorito"
		- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
	*/

	a := [][]string{{"Nome", "Sobrenome", "Hobby favorito"},
		{"Wagner", "Abrantes", "Filmes"},
		{"Beatriz", "Abrantes", "Ru Paul"},
		{"Maria", "Andrade", "Reclamar"},
	}

	for _, v := range a[0] {
		fmt.Printf("%v | ", v)
	}
	fmt.Print("\n")
	for _, v := range a[1] {
		fmt.Printf("%v | ", v)
	}
	fmt.Print("\n")
	for _, v := range a[2] {
		fmt.Printf("%v | ", v)
	}
	fmt.Print("\n")
	for _, v := range a[3] {
		fmt.Printf("%v | ", v)
	}

	// https://play.golang.org/p/dyTkWS2Hcre
}

func mapa() {
	/*
				- Crie um map com key tipo string e value tipo []string.
		    - Key deve conter nomes no formato sobrenome_nome
		    - Value deve conter os hobbies favoritos da pessoa
		- Demonstre todos esses valores e seus índices.
	*/
	a := map[string][]string{
		"Abrantes_Wagner":  []string{"Filmes", "Programmar"},
		"Abrantes_Beatriz": []string{"Ru Paul", "Beber Chá"},
		"Andrade_Maria":    []string{"Reclamar", "Igreja"},
	}

	//fmt.Println(a)

	/*
		- Utilizando o exercício anterior, adicione uma entrada ao map
		 e demonstre o map inteiro utilizando range.
	*/

	a["Alencar_Guilherme"] = []string{"Canceroso", "Real Madrid"}

	// Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
	delete(a, "Abrantes_Beatriz")

	for k, v := range a {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}

	// https://play.golang.org/p/kEyJ4m_LStJ
	// https://play.golang.org/p/Hppoc3yMYUW
	// https://play.golang.org/p/cDz3DeMqo_T
}
