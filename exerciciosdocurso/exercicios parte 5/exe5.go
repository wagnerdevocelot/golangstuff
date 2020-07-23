package main

import "fmt"

func main() {
	//sabores()
	//mapaStruct()
	//carros()
	//strucAnnonymous()
}

func sabores() {
	/*
				- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
		    - Nome
		    - Sobrenome
		    - Sabores favoritos de sorvete
		- Crie dois valores do tipo "pessoa" e
		 demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
	*/
	type pessoa struct {
		nome      string
		sobrenome string
		sabores   []string
	}

	pessoa1 := pessoa{
		nome:      "Wagner",
		sobrenome: "Abrantes",
		sabores:   []string{"Morango", "Chocolate", "Baunilha"},
	}

	pessoa2 := pessoa{
		nome:      "Beatriz",
		sobrenome: "Abrantes",
		sabores:   []string{"Pipoca", "Leite Condensado", "Milho"},
	}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.sabores {
		fmt.Println("\t", v)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.sabores {
		fmt.Println("\t", v)
	}
	// https://play.golang.org/p/urAK-KwUKte
}

func mapaStruct() {
	/*
				- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map,
				utilizando os sobrenomes como key.
		- Demonstre os valores do map utilizando range.
		- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
	*/
	type pessoa struct {
		nome      string
		sobrenome string
		sabores   []string
	}

	map1 := make(map[string]pessoa)

	map1["Abrantes"] = pessoa{
		nome:      "Wagner",
		sobrenome: "Abrantes",
		sabores:   []string{"Morango", "Chocolate", "Baunilha"},
	}

	for _, v := range map1 {
		fmt.Println("Oi, meu nome é", v.nome, "e meus sabores de sorvete favoritos são: ")
		for _, s := range v.sabores {
			fmt.Print("-", s, "-")
		}
		fmt.Print("\n")
	}
	// https://play.golang.org/p/zNaNCmJUJmS
}

func carros() {

	/*
				- Crie um novo tipo: veículo
		    - O tipo subjacente deve ser struct
		    - Deve conter os campos: portas, cor
		- Crie dois novos tipos: caminhonete e sedan
		    - Os tipos subjacentes devem ser struct
		    - Ambos devem conter "veículo" como struct embutido
		    - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
		    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
		- Usando os structs veículo, caminhonete e sedan:
		    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
		    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
		- Demonstre estes valores.
		- Demonstre um único campo de cada um dos dois.
	*/

	type veículo struct {
		portas int
		cor    string
	}

	type caminhonete struct {
		veículo
		traçãoNasQuatro bool
	}

	type sedan struct {
		veículo
		modeloLuxo bool
	}

	hilux := caminhonete{
		veículo: veículo{
			portas: 4,
			cor:    "Branco",
		},
		traçãoNasQuatro: true,
	}

	corolla := sedan{
		veículo: veículo{
			portas: 4,
			cor:    "Preto",
		},
		modeloLuxo: false,
	}

	fmt.Println(corolla.portas)
	fmt.Println(hilux.traçãoNasQuatro)

	// https://play.golang.org/p/5eHhTgM_vXH

}

func strucAnnonymous() {
	/*
			- Crie e use um struct anônimo.
		- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
	*/

	solidSnake := struct {
		equip  map[string]int
		weapon []string
	}{
		equip:  map[string]int{"Cigarette": 10, "Bandana": 1},
		weapon: []string{"Tranquilizer", "Jammer"},
	}

	fmt.Println(solidSnake)

	// https://play.golang.org/p/t-yFSr7NUIL
}
