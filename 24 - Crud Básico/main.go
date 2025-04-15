package main

import (
	"crud/servidor"
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

	rounter.HandleFunc("/usuarios", servidor.CriarUsuarios).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":5000", rounter))
}
