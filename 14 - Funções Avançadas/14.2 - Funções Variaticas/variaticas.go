package main

import (
	"fmt"
)

func somarNumeros(numeros ...int) int {
	total := 0
	for _, valor := range numeros {
		total += valor
	}
	return total
}

func exibirNumeros(texto string, numeros ...int) {
	for _, valor := range numeros {
		fmt.Println(texto, valor)
	}
	
}

func main () {
	fmt.Println("Funções Variaticas")

	fmt.Println(somarNumeros(1,4,5))

	exibirNumeros("Este numero é", 1,6,9)
}