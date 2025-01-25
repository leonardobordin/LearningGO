package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Ola Mundo!")
	escrever("Segundo escrever")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
