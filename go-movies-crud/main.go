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
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firestname string `json:"firstname"`
	Lastname   string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")
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
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "8382278", Title: "Vinnaithaandi Varuvaayaa", Director: &Director{Firestname: "Gautham ", Lastname: "Vasudev Menon"}})
	movies = append(movies, Movie{ID: "2", Isbn: "7481190", Title: "Raavanan", Director: &Director{Firestname: "Mani ", Lastname: "Ratnam"}})
	movies = append(movies, Movie{ID: "3", Isbn: "1483777", Title: "Subramaniapuram", Director: &Director{Firestname: "Sasi", Lastname: "Kumar"}})
	movies = append(movies, Movie{ID: "4", Isbn: "1053816", Title: "Mozhi", Director: &Director{Firestname: "Radha", Lastname: "Mohan"}})
	movies = append(movies, Movie{ID: "5", Isbn: "1274293", Title: "Vaaranam Aayiram", Director: &Director{Firestname: "Gautham ", Lastname: "Vasudev Menon"}})
	movies = append(movies, Movie{ID: "6", Isbn: "10189514", Title: "Soorarai Pottru", Director: &Director{Firestname: "Sudha", Lastname: "Kongara"}})
	movies = append(movies, Movie{ID: "7", Isbn: "0367495", Title: "Anbe Sivam", Director: &Director{Firestname: "Sunder", Lastname: "C"}})
	movies = append(movies, Movie{ID: "8", Isbn: "8114980", Title: "Pariyerum Perumal", Director: &Director{Firestname: " Mari", Lastname: "Selvaraj"}})
	movies = append(movies, Movie{ID: "9", Isbn: "6148156", Title: "Vikram Vedha", Director: &Director{Firestname: "Gayathri", Lastname: "K"}})
	movies = append(movies, Movie{ID: "10", Isbn: "3569782", Title: "Jigarthanda", Director: &Director{Firestname: "Karthik", Lastname: "Subbaraj"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting serverat port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
