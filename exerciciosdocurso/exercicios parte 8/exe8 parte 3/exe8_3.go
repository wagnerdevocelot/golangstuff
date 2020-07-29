/*
- Partindo do código abaixo, utilize NewEncoder() e Encode() para enviar as informações como JSON para Stdout.
    - https://play.golang.org/p/BVRZTdlUZ_
- Desafio: descubra o que é, e utilize, method chaining para conectar os dois métodos acima.
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	// slice de de structs
	users := []user{u1, u2, u3}

	//fmt.Println(users)

	// your code goes here

	// importando newenconder selecionando o metodo de saida
	saida := json.NewEncoder(os.Stdout)

	// ou err := json.NewEncoder(os.Stdout).Encode(users)
	// nessa segunda forma se cria uma cadeia de metodos interligados, como o Encode pertence ao NewEncoder
	// você pode chamar ele logo após sem intermediar através de uma variável

	// usando a importação de newencoder para imprimir o slice de structs no formato json no stdout(terminal)
	err := saida.Encode(users)
	if err != nil {
		fmt.Println(err)
	}

}
