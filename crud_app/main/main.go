package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http" //create a http server in Go
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json: "id"`
	Isbn     string    `json: "isbn"`
	Title    string    `json: "title"`
	Director *Director `json: "director"`
}

type Director struct {
	FirstName string `json: "firstname"`
	LastName  string `json: "lastname"`
}

var movies []Movie

func main() {

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "438",
		Title: "My 1st Moview",
		Director: &Director{
			FirstName: "Juan",
			LastName:  "Camacho",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "538",
		Title: "My 2nd Moview",
		Director: &Director{
			FirstName: "Juan",
			LastName:  "Camacho",
		},
	})

	r := mux.NewRouter()                                  //init instance for API
	r.HandleFunc("/movies", getMovies).Methods("GET")     //Path route to handle the API call , and the function with the logic. REST Method
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET") //this receives the ID to show the info related with the ID
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	log.Println("Server started ")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params passed by the API Request
	for i, item := range movies {
		if item.ID == params["id"] {
			//delete item
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies) //show the remaining movies on the slice
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params passed by the API Request
	for _, item := range movies {
		if item.ID == params["id"] {
			//get the item and encode in JSON Format
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) //Dcode the params from the request body and matches it as the JSON on the struct movies
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie) //Show the new movie added
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params passed by the API Request
	for i, item := range movies {
		if item.ID == params["id"] {
			var movie Movie
			//delete item
			movies = append(movies[:i], movies[i+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			json.NewEncoder(w).Encode(movie) //Show the new movie modified
			return
		}
	}
}
