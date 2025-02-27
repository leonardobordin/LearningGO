package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func(){
		escrever("Primeira Parte!")
		waitGroup.Done() // pega o Add que era 2 -1 
	}()
	
	go func(){ 
		escrever("Nos Sabemos!") 
		waitGroup.Done() // Pega o Add que era 2 -1
	}()

	waitGroup.Wait() // Espera chegar a 0
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
