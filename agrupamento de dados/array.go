package main

import "fmt"

// definindo um array de 5 posições de numeros inteiros
var x [5]int

func main() {
	//slaices()
	//sliceComFor()
	//arreis()
	//fatioPassou()
	//makingAMake()
	//slice2d()
	//mapa()
}

func mapa() {
	// mapas são estruturas de dados chave e valor e possuem mais de um tipo
	amigo := map[string]int{
		"Wagner":  11954560702,
		"Beatriz": 11954346676,
	}

	// adicionando novos itens ao map
	amigo["novo"] = 11976766688

	// acessando o valor da chave wagner
	fmt.Print(amigo["Wagner\n"])
	// ===> 11954560702
	fmt.Print(amigo["novo\n"])
	// ===> 11976766688

	// iterando com mapas
	for key, value := range amigo {
		fmt.Println(key, value)
	}

	// atraves da função delete passando a chave do valor deletamos o item desejado
	delete(amigo, "novo")

}

func slaices() {
	// não precisa definir o tamanho do slice pois a alocação é dinamica
	slice := []int{1, 2, 3}

	fmt.Println(slice)
	// ===> [1,2,3]

	// slices são mais flexiveis e você pode inserir novos valores
	slice = append(slice, 9)

	fmt.Println(slice)
	// ===> [1,2,3,9]

	// atribuindo um novo valor ao primeiro indice
	slice[0] = 1337
	fmt.Println(slice)
	// ===> [1337,2,3,9]

}

func sliceComFor() {
	// definição de um slice com string
	slice := []string{"Duke Nuken", "Quake", "Alex Kid"}

	// iteração do slice com o for usando o range(alcance) como parametro para acessar os valores do slice
	// nesse caso o range funciona como um contador num for normal com o seu valor pré definido pelo tamanho
	// do slice.
	for i, valor := range slice {
		fmt.Println("No indice", i, "temos o valor", valor)
		// ===> No indice 0 temos o valor Duke Nuken
		// ===>	No indice 1 temos o valor Quake
		// ===>	No indice 2 temos o valor Alex Kid
	}

}

func fatioPassou() {
	// um slice de strings normal
	sabores := []string{"Atum", "Portuguesa", "Abacaxi", "Bahiana", "Escarola"}
	// Indices   ======== [0] ====== [1] ======= [2] ====== [3] ======= [4] =======

	// fatia apartir do indice 1(Portuguesa) até o indice 4(Escarola)
	// apesar da fatia pedir o numero 4 isso quer dizer que esse é o limite da fatia, então seria
	// me ve as fatias ue vão do 1 até onde inicia a fatia 4
	fatias := sabores[1:4]

	fmt.Println(fatias)
	// ===> [Portuguesa Abacaxi Bahiana]

	// deletando todos os sabores entre atum e escarola e devolvendo um alice com esses dois sabores
	fatias = append(sabores[:1], sabores[4:]...)

	fmt.Println(fatias)
	// ===> [Atum Escarola]

}

func makingAMake() {
	// o make entrega uma forma mais performatica de trabalhar com slices sem precisar atribuir valores
	// definindo os indices e futuros indices, nesse caso ele tem 5 indices
	// mas se eu quiser utilzar os 10 eu posso utilizar o metodo append
	slice := make([]int, 5, 10)

	// atribuindo valores ao slice
	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	fmt.Println(slice)
	// ===> [1 2 3 4 5]

	// caso use o append mais alem do limite de 10 do slice ele cria um novo slice com o dobro de capacidade
}

func slice2d() {
	// slice multidimensional funciona basicamente com os mesmos conceitos dos arrays multidimensionais
	// são slices dentro de slices em que podemos usar dois indices, uma para linha e um para coluna
	ss := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	// para acessar slices 2D é meio que um trabalho de sniper vc precisa acertar os indices para
	// selecionar os itens corretos, nesse caso estamos printando 3 x 9
	fmt.Println(ss[0][2] * ss[2][2])
	// ===> 27
}

func arreis() {
	// definindo um array ja com seus itens
	y := [6]int{1, 2, 3, 4, 5, 6}
	// acessando e atribuindo valores aos indices do array
	x[0] = 5
	x[1] = 23
	// acessando pedaços especificos do array
	fmt.Println(x[0], x[1])
	// ===> 5, 23

	// acessando o array completo
	fmt.Println(x, y)
	// ===> x [5,23,0,0,0]
	// ===> y [1,2,3,4,5,6]

	// função leghth retorna a quantidade de indices de um array
	fmt.Println(len(x))
	// ===> 5
	// depois que um array é inicializado com o numero de indices não é possivel adicionar novos indices
	// só é possivel alterar os existentes
}
