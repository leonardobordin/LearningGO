package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
}

func (p pessoa) exibirNome() {
	fmt.Println(p.nome)
}

func (p *pessoa) acrescentarAnos(anos uint8) {
	p.idade+=anos
}

func main() {
	pessoa1 := pessoa{"Leonardo","Silva",26}
	pessoa1.exibirNome()
	pessoa1.acrescentarAnos(3)
	fmt.Println(pessoa1.idade)
}