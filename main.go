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

type Movie struct {
	Id       string    `json:"id"` // Field names must start with Cap to be exported in Go
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			w.WriteHeader(200)
			return
		}
	}
	w.WriteHeader(404)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// movie := []Movie{}
	for _, item := range movies {
		if item.Id == params["id"] {
			// fmt.Println(item)
			json.NewEncoder(w).Encode(item)
			w.WriteHeader(200)
			return
		}
	}
	w.WriteHeader(404)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	id, _ := strconv.Atoi(movies[len(movies)-1].Id)
	movie.Id = strconv.Itoa(id + 1)
	movie.Isbn = strconv.Itoa((rand.Intn(99999999)))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	for index, item := range movies {
		if item.Id == params["id"] {
			movie.Id = item.Id
			movie.Isbn = item.Isbn
			moviesEnd := movies[index+1:]
			movies = append(movies[:index], movie)
			movies = append(movies, moviesEnd...)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	w.WriteHeader(404)

}

func main() {

	movies = append(movies, Movie{Id: "1", Isbn: "109833", Title: "Worst: The Movie", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{Id: "2", Isbn: "831479", Title: "What is Entertainment", Director: &Director{FirstName: "Jane", LastName: "Eod"}})

	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", r))
}
