package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// CRUD - Create, Read, Update, Delete

	// CREATE - POST
	// READ - GET
	// UPDATE - PUT
	// DELETE - DELETE

	rounter := mux.NewRouter()

	rounter.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	rounter.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	rounter.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)

	fmt.Println("Escutando na porta 5000")

	log.Fatal(http.ListenAndServe(":5000", rounter))
}
