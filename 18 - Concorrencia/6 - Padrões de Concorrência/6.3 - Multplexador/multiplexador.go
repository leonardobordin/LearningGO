package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexador(escrever("Olá Mundo"), escrever("Programando em GO"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexador(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case texto := <-canalDeEntrada1:
				canalDeSaida <- texto
			case texto := <-canalDeEntrada2:
				canalDeSaida <- texto
			}
		}
	}()

	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Esse é o texto que está sendo passado: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
