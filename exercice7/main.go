package main

import (
	"fmt"
	"sync"
)

func createGoroutines() chan int {
	done := make(chan int, 10)

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			done <- i
		}()
	}

	// Fermer le channel une fois que toutes les goroutines ont terminé
	go func() {
		wg.Wait()
		close(done)
	}()

	return done
}

func main() {
	ch := createGoroutines()
	for i := range ch {
		fmt.Println(i)
	}
}
