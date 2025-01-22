package main

import (
	"fmt"
)

func closure() func() {
	texto := "Dentro da função clousure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao

}

func main(){
	texto := "Dentro da Função MAIN"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}