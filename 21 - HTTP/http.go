package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func main() {
	fmt.Println("HTTP")

	http.HandleFunc("/home",home)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
