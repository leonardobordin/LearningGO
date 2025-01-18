package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	idade int8
	endereco endereco
}

type endereco struct {
	rua string
	numero int16
}

func main() {
	fmt.Println("Arquivo Structs")

	var p pessoa
	p.nome = "Leonardo"
	p.idade = 26
	fmt.Println(p)

	pessoa2 := pessoa{"Giovanna", 27, endereco{"Rua do Pagode", 430}}
	fmt.Println(pessoa2)

	pessoa3 := pessoa{nome: "Ricardo"}
	fmt.Println(pessoa3)

	enderecoPessoa4 := endereco{"Avenido Doidera",234}
	pessoa4 := pessoa{"Ricardo", 24, enderecoPessoa4}
	fmt.Println(pessoa4.nome + " mora na rua " + pessoa4.endereco.rua)


}