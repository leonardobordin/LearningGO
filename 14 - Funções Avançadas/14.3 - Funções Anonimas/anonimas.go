package main

import (
	"fmt"
)

func main(){

	retorno := func(texto string, numero int) string {
		return fmt.Sprintf("Texto passado por parâmetro -> %s é %d",texto, numero)
	}("Leozin", 10)

	fmt.Println(retorno)
}