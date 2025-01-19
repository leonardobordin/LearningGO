package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estruturas de Controle")

	var numero int64 = 15

	if numero > 15 {
		fmt.Println("Numero maior que 15")
	} else if numero > 10 {
		fmt.Println("Numero menor que 15 porem maior que 10")
	} else {
		fmt.Println("Numero menor que 10")
	}

	numero++

	if outroNumero := numero; outroNumero > 15 {
		fmt.Println("Numero maior que 15")
	} else if numero > 10 {
		fmt.Println("Numero entre 10 e 15")
	} else {
		fmt.Println("Numero menor que 10")
	}
}
