package main

import (
	"fmt"
	"sync"
)

func main() {
	done := make(chan int, 10)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 1; i <= 10; i++ {
			done <- i
		}
		defer close(done)
		wg.Done()
	}()

	// Fermer le channel une fois que toutes les goroutines ont terminé
	go func() {
		for i := range done {
			fmt.Printf("%d\n", i)
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("\nFin du programme")

}
