package repository

import (
	"api-pokemon/entity"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
	//api_pokemon "api-pokemon"
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

func TestPokemonRepositoryImpl_Insert(t *testing.T) {
	pokemonRepository := NewPokemonRepository(GetConnection())
	ctx := context.Background()

	pokemon1 := entity.Pokemon{
		Name:    "Giratina",
		Type:    "Dragon & Dark",
		Species: "unknown",
	}

	result, err := pokemonRepository.Insert(ctx, pokemon1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestPokemonRepositoryImpl_FindAll(t *testing.T) {
	pokemonRepository := NewPokemonRepository(GetConnection())
	ctx := context.Background()

	result, err := pokemonRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, value := range result {
		fmt.Println(value.Id)
		fmt.Println(value.Name)
		fmt.Println(value.Type)
		fmt.Println(value.Species)
		fmt.Println("")

	}
}

func TestPokemonRepositoryImpl_FindById(t *testing.T) {
	pokemonRepository := NewPokemonRepository(GetConnection())
	ctx := context.Background()

	PokeId := 1
	result, err := pokemonRepository.FindById(ctx, int32(PokeId))
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestPokemonRepositoryImpl_UpdateById(t *testing.T) {
	pokemonRepository := NewPokemonRepository(GetConnection())
	ctx := context.Background()

	PokeId := 1
	updatePoke1 := entity.Pokemon{
		Name:    "Giratina",
		Type:    "Dragon & Dark",
		Species: "Mytical Zoan ",
	}

	result, err := pokemonRepository.UpdateById(ctx, updatePoke1, int32(PokeId))
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

}

func TestPokemonRepositoryImpl_DeleteById(t *testing.T) {
	pokemonRepository := NewPokemonRepository(GetConnection())
	ctx := context.Background()

	PokeId := 3
	err := pokemonRepository.DeleteById(ctx, int32(PokeId))
	if err != nil {
		panic(err)
	}

	fmt.Println("Congrats success delete pokemon Id", PokeId)
}
