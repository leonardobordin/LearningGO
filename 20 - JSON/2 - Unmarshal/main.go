package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type pessoa struct {
	Nome   string `json:"nome"`
	Genero string `json:"genero"`
}

func main() {
	fmt.Println("JSON Unmarshal")

	pessoaEmJSON := `{"nome":"Giovanna", "genero":"Feminino"}`
	fmt.Println(pessoaEmJSON)
	var p pessoa
	if erro := json.Unmarshal([]byte(pessoaEmJSON), &p); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(p)

	pessoa2EmJSON := `{"nome":"Leonardo", "genero":"Masculino"}`
	p2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(pessoa2EmJSON), &p2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(p2)

}
