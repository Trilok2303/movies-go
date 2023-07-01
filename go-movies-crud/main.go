package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var mov movie
	_ = json.NewDecoder(r.Body).Decode(&mov)
	mov.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, mov)
	json.NewEncoder(w).Encode(mov)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var mov movie
			_ = json.NewDecoder(r.Body).Decode(&mov)
			mov.ID = strconv.Itoa(rand.Intn(100000000))
			movies = append(movies, mov)
			json.NewEncoder(w).Encode(mov)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()
	movies = append(movies, movie{ID: "1", Isbn: "438227", Title: "Movie one", Director: &Director{FirstName: "Trilok", LastName: "Jain"}})
	movies = append(movies, movie{ID: "2", Isbn: "438228", Title: "Movie two", Director: &Director{FirstName: "Trilok1", LastName: "Jain1"}})
	movies = append(movies, movie{ID: "3", Isbn: "438229", Title: "Movie three", Director: &Director{FirstName: "Trilok2", LastName: "Jain2"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	fmt.Printf("Startng server port Number is 8116\n")
	log.Fatal(http.ListenAndServe(":8116", r))
}
