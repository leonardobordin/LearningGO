package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type pessoa struct {
	Nome   string `json:"nome"` // Se colocar um "-" ao invés da tag "nome" ele ignora o campo na conversão
	Genero string `json:"genero"`
	Idade  uint   `json:"idade"`
}

func main() {
	fmt.Println("JSON Marshal")

	p1 := pessoa{"Leonardo", "Masculino", 26}

	fmt.Println(p1)

	pessoaEmJSON, erro := json.Marshal(p1)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(pessoaEmJSON)
	fmt.Println(bytes.NewBuffer(pessoaEmJSON))

	p2 := map[string]string{
		"nome":   "Giovanna",
		"genero": "Feminino",
	}

	pessoa2EmJSON, erro := json.Marshal(p2)
	fmt.Println(pessoa2EmJSON)
	fmt.Println(bytes.NewBuffer(pessoa2EmJSON))
}
