package main

import (
	"fmt"
	//"time"
)

func main() {
	fmt.Println("Loops")

	i := 0
	for i < 10 {
		i++
		fmt.Println("Em qual loop estou: ", i)
		//time.Sleep(time.Second)
	}
	
	fmt.Println("---------------------------------------------")
 
	for j := 0; j < 10 ; j++ {
		fmt.Println("Loop Definindo Em Uma linha. Posição: ",j)
	}

	fmt.Println("---------------------------------------------")

	array := [5]string{"João","Matheus","Leonardo","Renato","Jhon"}
	for indice, valor := range(array) {
		fmt.Println(indice, valor)
	}

	for _, valor := range(array) {
		fmt.Println(valor)
	}

	fmt.Println("---------------------------------------------")

	map1 := map[string]string {
		"nome": "Giovanna",
		"sobrenome": "Benichel",
	}

	for chave, valor := range(map1) {
		fmt.Println(chave, valor)
	}
}
