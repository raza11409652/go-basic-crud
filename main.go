package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movie []Movie

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.Host)
}

func GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)

	movie = append(movie, Movie{ID: "1", Name: "Movie 1", Isbn: "123456789", Title: "Movie Title 1", Director: &Director{FirstName: "First name", LastName: "Last name"}})
	movie = append(movie, Movie{ID: "2", Name: "Movie 2", Isbn: "123456790", Title: "Movie Title 2", Director: &Director{FirstName: "First name", LastName: "Last name"}})
	r.HandleFunc("/movies", GetMoviesHandler)
	fmt.Print("Starting server at port 2021")
	http.ListenAndServe(":2021", r)
}
