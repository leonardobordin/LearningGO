package main

import (
	"fmt"
)

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	fmt.Println("Funções com ponteiro")
	numero := 20;
	fmt.Println(numero)
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)

	numeroPonteiro := 40
	fmt.Println(numeroPonteiro)
	inverterSinalComPonteiro(&numeroPonteiro)
	fmt.Println(numeroPonteiro)
}