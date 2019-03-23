package main

import (
	"fmt"
	"log"
	"net/http"
	"encodig/json"
	"github.com/gorilla/mux"
)

//Index metod
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo desde mi servidor web con GO")
}

//MovieList metod
func MovieList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listado de Peliculas")

	movies := Movies{
		Movie{"Sin Limites", 2013, "Desconocido"},
		Movie{"Busqueda Implacable", 2009, "Desconocido"}
	}

	json.NewEncoder(w).Encode(movies)
}

//MovieShow metod
func MovieShow(w http.ResponseWriter, r *http.Request) {
	//Obtengo los parametros de la url
	params := mux.Vars(r)
	//Obtengo el parametro que necesito en particular
	movieID := params["id"]

	fmt.Fprintf(w, "Cargaste la pelicula numero %s", movieID)
}