package main

import (
	"fmt"
)

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	// 1 1 2 3 5 8
	posicao := uint(6)
	fmt.Println(fibonacci(posicao))
}
