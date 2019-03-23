package main

import (
	"log"
	"net/http"
)

func main() {
	//Para cuando no utilizamos libreria de ruteo
	//	http.HandleFunc("/", func(wr http.ResponseWriter, r *http.Request) {
	//		fmt.Fprintf(wr, "Hola Mundo desde mi Servidor Web con Go")
	//	})
	router := NewRouter()

	server := http.ListenAndServe(":8080", router)
	log.Fatal(server)
}
