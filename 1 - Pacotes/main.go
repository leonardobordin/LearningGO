package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Olá Mundo !")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("kdiewkdwe")
	fmt.Println(erro)
}
