package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("avenida Vila dos Bobos")
	fmt.Println(tipoEndereco)
}
