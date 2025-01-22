package main

import (
	"fmt"
)

var n int

func init() {
	fmt.Println("Função init unica que executa antes da main")
	n = 10
}

func main() {
	fmt.Println("Dentro da função main")
	fmt.Println(n)
}
