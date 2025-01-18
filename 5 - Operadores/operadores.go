package main

import (
	"fmt"
)

func main() {
	// Operadores Aritimeticos
	soma := 10 + 5
	subtracao := 5 - 1
	divisao := 10 / 5
	multiplicacao := 2 * 2
	restoDivisao := 3 % 2 	
	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)
	var numero1 int16 = 20
	var numero2 int16 = 40
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	// Fim dos Operadores Aritimeticos

	// Operadores de Atribuicao
	var variavel1 string = "String"
	variavel2 := "String"
	fmt.Println(variavel1, variavel2)
	// Fim do Operadores de Atribuicao

	//Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(3 >= 4)
	fmt.Println(9 == 9)
	fmt.Println(4 <= 3)
	fmt.Println(6 != 8)
	fmt.Println("-------------------")
	// Fim do Operadores Relacionais

	// Operadores Lógicos
	verdadeiro, falso := true, false
	fmt.Print(verdadeiro && falso)
	fmt.Print(verdadeiro || falso)
	fmt.Print(!verdadeiro)
	// Fim dos Operadores Lógicos

	// Operadores Unarios
	var numero int8 = 5
	fmt.Println(numero)
	numero++ // numero = numero + 1
	fmt.Println(numero)
	numero-- // numero = numero - 1
	fmt.Println(numero)
	numero += 10 // numero = numero + 10
	fmt.Println(numero)
	numero -= 2 // numero = numero - 2
	fmt.Println(numero)
	numero *= 2 // numero = numero * 2
	fmt.Println(numero)
	// Fim do Operadores Unarios

	if numero > 5 {
		fmt.Println("Numero maior que 5")
	} else {
		fmt.Println("Numero menor que 5")
	}
}