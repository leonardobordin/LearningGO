package main

import (
	"fmt"
)

func funcao1(){
	fmt.Println("Escrevendo na função 1")
}

func funcao2(){
	fmt.Println("Escrevendo na função 2")
}

func alunoAprovado(n1,n2,n3 float32) bool {
	defer fmt.Println("Media calculado com sucesso")
	fmt.Println("Realizando calculo")

	media := (n1+n2+n3)/3

	if media >= 7 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Defer = Adiar")

	defer funcao1()
	funcao2()
	aprovado := alunoAprovado(1,9,7)
	fmt.Println("Aluno aprovado:", aprovado)
}