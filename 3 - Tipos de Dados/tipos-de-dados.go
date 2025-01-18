package main

import(
	"fmt"
	"errors"
)

func main() {
	// Inteiros
	var inteiro8 int8 = 123
	fmt.Println(inteiro8)
	var inteiro16 int16 = 12345
	fmt.Println(inteiro16)
	var inteiro32 int32 = 1234567890
	fmt.Println(inteiro32)
	var inteiro64 int64 = 1234567890123456789
	fmt.Println(inteiro64)

	// float
	var flutuante32 float32 = 12.12
	fmt.Println(flutuante32)
	var flutuante64 float64 = 12345.12
	fmt.Println(flutuante64)

	// Strings
	var texto1 string = "texto1"
	fmt.Println(texto1)
	texto2 := "texto2"
	fmt.Println(texto2)

	char := 'A'
	fmt.Println(char)

	// Bool
	var verdadeiro1 bool = false
	verdadeiro2 := true
	fmt.Println(verdadeiro1, verdadeiro2)

	// Error
	var erroNulo error
	fmt.Println(erroNulo)
	var erro error = errors.New("Erro Interno")
	fmt.Println(erro)

}