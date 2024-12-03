package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type Movie struct{
		ID string `json:"id"`
		Isbn string `json:"isbn"`
		Title string `json:"title"`
		Director *Director `json:"director"`
}

type Director struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}


var movies [] Movie;

func main () {
		r := mux.NewRouter();

		r.HandleFunc("/movies", getMovies).Methods("GET");
		r.HandleFunc("/movie/{id}", getMovie).Methods("GET");
		r.HandleFunc("/movie", createMovie).Methods("POST");
		r.HandleFunc("/movie/{id}", updateMovie).Methods("PUT");
		r.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE");


		fmt.Printf("Starting server at port 8000\n")
		log.Fatal(http.ListenAndServe(":8000",r))
}