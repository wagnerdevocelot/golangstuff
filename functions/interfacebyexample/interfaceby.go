package main

import (
	"fmt"
	"math"
)

// tipo retangulo com dois valores float largura e altura
type retangulo struct {
	largura, altura float64
}

// tipo circulo que tem o tipo raio que também é um float
type circulo struct {
	raio float64
}

// tipo interface recebe dois metodos area e perimetro ambos float
type geometria interface {
	area() float64
	perimetro() float64
}

// metodo area retorna o valor da altura X largura nesse caso o receiver é um retangulo
func (r retangulo) area() float64 {
	return r.largura * r.altura
}

// metodo retorna o valor de (largura x 2) + (altura x 2) receiver é um retangulo
func (r retangulo) perimetro() float64 {
	return 2*r.largura + 2*r.altura
}

// metodo area usa o numero Pi = 3,14 x raio(ao quadrado) dessa vez calcurando o circulo como receiver
// interessante ressaltar que é o mesmo metodo fazendo uma coisa diferente (polimorfismo)
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}

// metodo perimetro retorna (3,14 x 2) x raio tendo o circulo como receiver
// interessante ressaltar que é o mesmo metodo fazendo uma coisa diferente (polimorfismo)
func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// função medidas recebe a interface geomatria como parametro
func medidas(g geometria) {

	fmt.Println(g)             // printa os valores do parametro, seja ele o circulo ou o retangulo
	fmt.Println(g.area())      // o mesmo agora usando o metodo area
	fmt.Println(g.perimetro()) // e agora usando o metodo perimetro
}

func main() {
	r := retangulo{largura: 3, altura: 4} // struct de retangulo
	c := circulo{raio: 5}                 // struct de circulo

	medidas(r) // a função medida printando os valores, area e perimetro do retangulo
	medidas(c) // a função medida printando os valores, area e perimetro do circulo
}
