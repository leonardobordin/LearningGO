package main

import (
	"fmt"
)

func main() {
	canal := make(chan string,2)
	canal <- "Olá Mundo"
	canal <- "Estou programando"
	//go atribuirValores("Olá Mundo!", canal)

	mensagem := <-canal
	mensagem2 := <- canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}

//func atribuirValores(texto string, canal chan string){
//	canal <- texto
//	close(canal)
//}