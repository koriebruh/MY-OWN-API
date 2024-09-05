package main

import (
	"api-pokemon/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// <-- we use Mysql driver and Gorilla mux

func main() {

	db := handlers.GetConnection()
	defer db.Close()

	r := mux.NewRouter()
	r.HandleFunc("/pokemons", handlers.GetPokemons).Methods("GET")           // Get all Pokemons
	r.HandleFunc("/pokemons/{id}", handlers.GetPokemonById).Methods("GET")   // Get Pokemon by ID
	r.HandleFunc("/pokemons", handlers.CreatePokemon).Methods("POST")        // Create new Pokemon
	r.HandleFunc("/pokemons/{id}", handlers.UpdatePokemon).Methods("PUT")    // Update Pokemon by ID
	r.HandleFunc("/pokemons/{id}", handlers.DeletePokemon).Methods("DELETE") // Delete Pokemon by ID

	log.Println("Server listening on port :8080")
	http.ListenAndServe("localhost:8080", r)

}
