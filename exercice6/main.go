package main

import (
	"fmt"
	"time"
)

func oneToTen() {

	for i := 0; i <= 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}

func aToZ() {

	for i := 'A'; i <= 'Z'; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(string(i))
	}
}

func main() {
	done := make(chan bool)
	go func() {
		oneToTen()
		done <- true
	}()
	go func() {
		go aToZ()
		done <- true
	}()

	<-done
	<-done

}
