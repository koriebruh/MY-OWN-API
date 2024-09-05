package handlers

import (
	"api-pokemon/entity"
	"api-pokemon/repository"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:korie123@tcp(127.0.0.1:3306)/api_pokemon?parseTime=true")
	// <-- ?parseTime=true , di tambahkan itu jika nanti parse ke golang otomatis jadi bisa tipe data Time.Time
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxIdleConns(10)                  //<-- berapa minimal koneksi saat pertama jalankan appnya
	db.SetMaxOpenConns(100)                 //<-- maxsimal koneksi,
	db.SetConnMaxIdleTime(5 * time.Minute)  //<-- berapa lama koneksi yang sudah tidak digunakan akan dihapus
	db.SetConnMaxLifetime(60 * time.Minute) //<-- seberapa lama koneksi boleh digunakan

	return db
}

func GetPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// <-- meelakukan query sql
	pokemonRepository := repository.NewPokemonRepository(GetConnection())
	ctx := context.Background()
	result, err := pokemonRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	// <-- endcode mengirim data hasul query fromat json
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, fmt.Sprintf("Encoder() error: %v", err), http.StatusInternalServerError)
	}

}

func GetPokemonById(w http.ResponseWriter, r *http.Request) {
	//<-- take request param
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"]) //<-- mengubah param dari web jadi int
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// <-- query
	ctx := context.Background()
	pokemonRepository := repository.NewPokemonRepository(GetConnection())
	result, err := pokemonRepository.FindById(ctx, int32(id))
	if err != nil {
		panic(err)
	}

	// <-- endcode mengirim data hasul query fromat json
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, fmt.Sprintf("Encoder() error: %v", err), http.StatusInternalServerError)
	}

}

func CreatePokemon(w http.ResponseWriter, r *http.Request) {
	//<-- take request from json to struct
	w.Header().Set("Content-Type", "application/json")
	var pokemonAdd entity.Pokemon
	err := json.NewDecoder(r.Body).Decode(&pokemonAdd)

	//<-- execute query
	pokemonRepository := repository.NewPokemonRepository(GetConnection())
	ctx := context.Background()
	result, err := pokemonRepository.Insert(ctx, pokemonAdd)
	if err != nil {
		panic(err)
	}

	// <-- endcode mengirim data hasil query fromat json
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, fmt.Sprintf("Encoder() error: %v", err), http.StatusInternalServerError)
	}
}

func UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	//<-- take request from json to struct
	var pokemonUpdate entity.Pokemon
	err := json.NewDecoder(r.Body).Decode(&pokemonUpdate)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	//<-- take request param
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"]) //<-- mengubah param dari web(mapString) jadi int
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	///////////////////////////////////////////////////////////////////////////////////
	// <-- query
	pokemonRepository := repository.NewPokemonRepository(GetConnection())
	ctx := context.Background()
	result, err := pokemonRepository.UpdateById(ctx, pokemonUpdate, int32(id))
	if err != nil {
		panic(err)
	}

	// <-- endcode mengirim data hasil query fromat json
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, fmt.Sprintf("Encoder() error: %v", err), http.StatusInternalServerError)
	}

}

func DeletePokemon(w http.ResponseWriter, r *http.Request) {
	//<-- take request param
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"]) //<-- mengubah param dari web(mapString) jadi int
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	// <-- query
	pokemonRepository := repository.NewPokemonRepository(GetConnection())
	ctx := context.Background()
	if err := pokemonRepository.DeleteById(ctx, int32(id)); err != nil {
		http.Error(w, "Param not found", http.StatusNotFound)
	}

	// Return a successful response with JSON output
	resp := Response{
		Status:  "success",
		Message: "Data berhasil dihapus",
	}
	json.NewEncoder(w).Encode(resp)
	w.WriteHeader(http.StatusOK) // or http.StatusNoConten

}

type Response struct {
	Status  string
	Message string
}
