package main

import (
	"fmt"
	"math/rand"
	"time"
)

var posseDeBola bool = true
var count1 int = 0
var count2 int = 0

func main() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i <= 90; i++ {
		time1()
		time2()
	}
	fmt.Printf("============Final de Jogo=========\n\n\n")
	fmt.Printf("====Corinthians %v x %v Palmeiras====\n", count1, count2)
}

func time1() {

	ataque := rand.Intn(100)
	if ataque > 70 {
		fmt.Println("Passou do meio campo, rumo ao ataque o Corinthians")
		if ataque > 97 {
			fmt.Println("É goooooooooooooooooooooool!!! Do Corinthians")
			count1++
		} else {
			posseDeBola = false
			fmt.Println("Perdeu a posse da bola o Corinthians")
		}
	} else {
		posseDeBola = false
		fmt.Println("Perde a posse da bola o Corinthians")
	}

}

func time2() {

	ataque := rand.Intn(100)
	if ataque > 70 {
		fmt.Println("Passou do meio campo, rumo ao ataque o Palmeiras")
		if ataque > 97 {
			fmt.Println("É goooooooooooooooooooooool!!! do Palmeiras")
			count2++
		} else {
			posseDeBola = false
			fmt.Println("Perdeu a posse da bola o Palmeiras")
		}
	} else {
		posseDeBola = false
		fmt.Println("Perde a posse da bola o Palmeiras")
	}

}

// https://play.golang.org/p/pEIUhbhXCgP
