package main

import (
	"fmt"
	"runtime"
	"sync"
)

// importando os metodos de wait group
var wg sync.WaitGroup

func main() {
	contador := 0
	totalGoRoutines := 1000

	// importando os metodos lock e unlock de mutex para que não haja race conditions durantes a
	// execussão das go routines
	var mu sync.Mutex

	// referencia ao total de go routines a serem executadas
	wg.Add(totalGoRoutines)

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			// Gosched cede a memoria da execução por um tempo limitado parecido com sleep()
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	// wait avisa main quando a execução das go routines terminam
	wg.Wait()

	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Valor do contador: ", contador)
}
