package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Create the class
type pokemon struct {
	// save variables as attributes
	name  string
	types string
	moves []string
	stats map[string]int
	hp    int
}

var pokemon1Attack string
var pokemon2Attack string
var index int

var pokemon1 = pokemon{
	name:  "Charizard",
	types: "Fire",
	moves: []string{"Flamethrower", "Fly", "Blast Burn", "Fire Punch"},
	stats: map[string]int{"ATTACK": 6, "DEFENSE": 5},
	hp:    10,
}

var pokemon2 = pokemon{
	name:  "Blastoise",
	types: "Water",
	moves: []string{"Water Gun", "Bubblebeam", "Hydro Pump", "Surf"},
	stats: map[string]int{"ATTACK": 6, "DEFENSE": 5},
	hp:    10,
}

func main() {
	//delayPrint("Its super effective...")
	fight(pokemon1, pokemon2)
}

// print one character at a time
func delayPrint(s ...string) {
	for _, c := range s {
		fmt.Print(string(c))
		time.Sleep(180 * time.Millisecond)
	}
}

func (p *pokemon) lessthanzero() {
	if p.hp < 0 {
		for p.hp != 0 {
			p.hp = p.hp + 1
		}
	}
}

func fight(one pokemon, two pokemon) {
	// Allow two pokemon to fight each other
	fmt.Println("-----POKEMONE BATTLE-----")
	fmt.Println(one.name)
	fmt.Println("TYPE/", one.types)
	fmt.Println("ATTACK/", one.stats["ATTACK"])
	fmt.Println("DEFENSE/", one.stats["DEFENSE"])
	fmt.Println("\nVS")
	fmt.Println(two.name)
	fmt.Println("TYPE/", two.types)
	fmt.Println("ATTACK/", two.stats["ATTACK"])
	fmt.Println("DEFENSE/", two.stats["DEFENSE"])

	time.Sleep(2 * time.Second)

	// consider types advantages
	version := []string{"Fire", "Water", "Grass"}

	for i, k := range version {
		if pokemon1.types == k {
			// both are same types
			if pokemon2.types == k {
				pokemon1Attack = "\nIts not very effective...\n"
				pokemon2Attack = "\nIts not very effective...\n"
			}

			// pokemon2 is STRONG
			if pokemon2.types == version[(i+1)%3] {
				pokemon2.stats["ATTACK"] *= 2
				pokemon2.stats["DEFENSE"] *= 2
				pokemon1.stats["ATTACK"] /= 2
				pokemon1.stats["DEFENSE"] /= 2
				pokemon1Attack = "\nIts not very effective...\n"
				pokemon2Attack = "\nIts super effective!\n"
			}

			// pokemon2 is WEAK
			if pokemon2.types == version[(i+2)%3] {
				pokemon1.stats["ATTACK"] *= 2
				pokemon1.stats["DEFENSE"] *= 2
				pokemon2.stats["ATTACK"] /= 2
				pokemon2.stats["DEFENSE"] /= 2
				pokemon1Attack = "\nIts super effective!\n"
				pokemon2Attack = "\nIts not very effective...\n"
			}
		}

	}

	// Actual fight
	// Continue while pokemon still have health
	for {

		// begins battle
		fmt.Printf("\n%v \t\tHLTH\t%v\n", pokemon1.name, pokemon1.hp)
		fmt.Printf("%v \t\tHLTH:\t%v\n\n", pokemon2.name, pokemon2.hp)

		// show moves
		fmt.Printf("Go %v!\n\n", pokemon1.name)
		for i, x := range pokemon1.moves {
			fmt.Printf("%v. %v\n", i+1, x)
		}

		// first turn battle
		fmt.Print("Pick a move: ")
		fmt.Scanln(&index)
		delayPrint("\n", pokemon1.name, " used ", pokemon1.moves[index-1], "!")
		time.Sleep(1 * time.Second)
		delayPrint(pokemon1Attack)

		// determine damage
		pokemon2.hp -= pokemon1.stats["ATTACK"]

		// adiciona mais 1 ao hp até ficar com 0 para que não termine com o hp negativo
		pokemon2.lessthanzero()
		// fim da implementação

		time.Sleep(1 * time.Second)
		fmt.Printf("\n%v \t\tHLTH\t%v\n", pokemon1.name, pokemon1.hp)
		fmt.Printf("%v \t\tHLTH:\t%v\n\n", pokemon2.name, pokemon2.hp)
		time.Sleep(1 * time.Second)

		// check if pokemon fainted
		if pokemon2.hp <= 0 {
			delayPrint("\n...", pokemon2.name, " fainted.")
			break
		}

		// pokemon2 turn
		// show moves
		fmt.Printf("Go %v!\n", pokemon2.name)
		for i, x := range pokemon2.moves {
			fmt.Printf("%v. %v\n", i+1, x)
		}

		// second turn battle
		fmt.Print("Pick a move: ")
		fmt.Scanln(&index)
		delayPrint("\n", pokemon2.name, " used ", pokemon2.moves[index-1], "!")
		time.Sleep(1 * time.Second)
		delayPrint(pokemon2Attack)

		// determine damage
		pokemon1.hp -= pokemon2.stats["ATTACK"]

		// adiciona mais 1 ao hp até ficar com 0 para que não termine com o hp negativo
		pokemon1.lessthanzero()

		time.Sleep(1 * time.Second)
		fmt.Printf("\n%v \t\tHLTH\t%v\n", pokemon1.name, pokemon1.hp)
		fmt.Printf("%v \t\tHLTH:\t%v\n\n", pokemon2.name, pokemon2.hp)
		time.Sleep(1 * time.Second)

		// check if pokemon1 fainted
		if pokemon1.hp <= 0 {

			delayPrint("\n...", pokemon1.name, " fainted.")
			break
		}

		if pokemon1.hp > 0 && pokemon2.hp > 0 {
			continue
		}

	}
	rand.Seed(time.Now().UnixNano())
	money := rand.Intn(3000)
	delayPrint("\nOpponent paid you: ")
	fmt.Printf("%vG\n", money)

}
