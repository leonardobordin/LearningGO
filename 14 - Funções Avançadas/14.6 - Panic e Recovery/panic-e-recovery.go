package main

import (
	"fmt"
)

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Tentativa de recuperar execução!")
	}
}

func alunoEstaAprovado(n1,n2 float32) bool {
	defer recuperarExecucao()
	fmt.Println("Realizando Calculos...")

	media := (n1+n2)/2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MÉDIA É 6")
}

func main () {
	fmt.Println("Panic e Recovery")

	fmt.Println(alunoEstaAprovado(6,6))
	fmt.Println("Finalizou")
}