package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8) {
	return n1 + n2, n1 - n2
}

func adicionarSobrenome(nome string) string {
	return nome + " " + "Bordin"
}

func main() {
	fmt.Println(somar(10, 5))

	resultadoSoma, resultadoSubtracao := calculos(10, 5)
	fmt.Println(resultadoSoma, resultadoSubtracao)

	fmt.Println(calculos(30, 2))

	resultadoSoma2, _ := calculos(60, 20) // o Underline ignora o retorno da função
	fmt.Println(resultadoSoma2)

	fmt.Print(adicionarSobrenome("Leonardo"))
}
