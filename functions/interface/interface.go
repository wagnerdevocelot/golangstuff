package main

import "fmt"

// struct
type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

// struct com struct aninhado
type dentista struct {
	pessoa
	dentesarrancados int
	salario          float64
}

// struct com struct aninhado
type arquiteto struct {
	pessoa
	tipodeconstrução string
	tamanhodaloucura string
}

// função com receiver
func (x dentista) oibomdia() {
	fmt.Println("Meu nome é: ", x.nome, "e eu ja arranquei", x.dentesarrancados, "dentes")
}

// função com receiver
func (y arquiteto) oibomdia() {
	fmt.Println("Meu nome é: ", y.nome, "e eu construo", y.tipodeconstrução)
}

// interface que recebe a função mas não recebe struct
type gente interface {
	oibomdia()
}

// função que recebe a interface como parametro
func serhumano(g gente) {
	g.oibomdia()
}

func main() {
	mrdente := dentista{
		pessoa: pessoa{
			nome:      "Mister Dente",
			sobrenome: "Da Silva",
			idade:     56,
		},
		dentesarrancados: 345,
		salario:          543456,
	}

	mrarquiteto := arquiteto{
		pessoa: pessoa{
			nome:      "Mister Arquiteto",
			sobrenome: "Mosby",
			idade:     56,
		},
		tipodeconstrução: "pontes",
		tamanhodaloucura: "grande",
	}

	// uso das funções atraves dos structs
	mrarquiteto.oibomdia()
	mrdente.oibomdia()

	fmt.Print("\n\n\n")

	// o uso da interface para usar metodos contidos nela atraves de structs como parametro
	serhumano(mrarquiteto)
	serhumano(mrdente)
}
