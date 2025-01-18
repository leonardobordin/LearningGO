package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ponteiros")
	
	var variavel1 int8 = 10
	variavel2 := variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	var variavel3 int8 = 40
	var ponteiro *int8
	fmt.Println(variavel3, ponteiro)

	ponteiro = &variavel3
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)
	variavel3 = 60
	fmt.Println(variavel3, ponteiro)
	fmt.Println(variavel3, *ponteiro)
}