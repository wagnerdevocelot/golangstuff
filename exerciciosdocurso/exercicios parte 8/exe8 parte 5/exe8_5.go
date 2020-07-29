/*
- Partindo do código abaixo, ordene os []user por idade e sobrenome.
    - https://play.golang.org/p/BVRZTdlUZ_
- Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.
*/

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type idade []user

func (a idade) Len() int           { return len(a) }
func (a idade) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a idade) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type lastname []user

func (a lastname) Len() int           { return len(a) }
func (a lastname) Less(i, j int) bool { return a[i].Last < a[j].Last }
func (a lastname) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

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

	users := []user{u1, u2, u3}

	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	sort.Sort(lastname(users))

	sort.Sort(idade(users))

	for _, v := range users {
		fmt.Println("Nome:", v.First, v.Last, "Idade:", v.Age)
		for _, j := range v.Sayings {
			fmt.Println("\t", j)
		}
	}

}
