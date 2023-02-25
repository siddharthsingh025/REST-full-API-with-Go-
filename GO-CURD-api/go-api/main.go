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
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

//function to get of movies in database

func getMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// function to delete particular movie in database
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

// function to get particular movie by its ID
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

// function to creat new movie record to databse
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode((&movie))
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

//function to update any particular movie record from existing records in the databse

// we here dont go for complex approch , we will just delete existing record or particular id and then re-add it to databse with new values geting from client request
func updateMovie(w http.ResponseWriter, r *http.Request) {
	// 1. set json content-type
	// 2. get params
	// 3. loop over existing records of movies , range
	// 4. delete the particular record with ID that you've sent
	// 5. add new movie record - the movie that we get in body of client request

	// 1
	w.Header().Set("Content-Type", "aplication/json")

	// 2
	params := mux.Vars(r)

	//3 , 4 , 5
	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie

			_ = json.NewDecoder(r.Body).Decode((&movie))
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {

	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie one", Director: &Director{Firstname: "Siddharth", Lastname: "Singh"}})
	movies = append(movies, Movie{ID: "2", Isbn: "435566", Title: "Movie two", Director: &Director{Firstname: "himanshu", Lastname: "Singh"}})
	movies = append(movies, Movie{ID: "3", Isbn: "435697", Title: "Movie three", Director: &Director{Firstname: "Muskan", Lastname: "Ladwal"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Print("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))

}
