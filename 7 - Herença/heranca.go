package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
}

type estudante struct{
	pessoa
	curso string
	escola string
}

func main() {
	p1 := pessoa{"Joao","Silva",16,165}
	fmt.Println(p1)

	e1 := estudante{p1, "Eletromecanica","Senai"}
	fmt.Println(e1.nome + " " + e1.sobrenome + " cursa " + e1.curso + " no " + e1.escola)
}