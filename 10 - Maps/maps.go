package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	usuario := map[string]string {
		"nome": "João",
		"sobrenome": "silva",
	}

	fmt.Println(usuario["sobrenome"])

	usuario2 := map[string]map[string]string{
		"dados_pessoais":{
			"nome": "Paulo",
			"sobrenome": "Silveira",
			"pai": "Carlos",
		},
		"endereco":{
			"rua": "Antonio Alves",
			"cidade": "Bauru",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "endereco")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"mes": "maio",
		"tipo": "gemeos",
	}
	fmt.Println(usuario2)

}