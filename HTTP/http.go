package main

import (
	"fmt"
	"log"
	"net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Raiz"))
}
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Servidor OK"))
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Carregando páginas de usuários..."))
}

func main() {
	fmt.Println("HTTP")

	http.HandleFunc("/", root)
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
