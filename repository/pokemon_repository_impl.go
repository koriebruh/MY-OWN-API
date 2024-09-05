package repository

import (
	"api-pokemon/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type pokemonRepositoryImpl struct {
	DB *sql.DB
}

func NewPokemonRepository(db *sql.DB) PokemonRepository {
	return pokemonRepositoryImpl{DB: db}
}

func (repository pokemonRepositoryImpl) Insert(ctx context.Context, pokemon entity.Pokemon) (entity.Pokemon, error) {
	sqlExec := "INSERT INTO pokemons (name,type,species) VALUES (?,?,?)"
	result, err := repository.DB.ExecContext(ctx, sqlExec, pokemon.Name, pokemon.Type, pokemon.Species)
	if err != nil {
		return pokemon, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return pokemon, err
	}

	pokemon.Id = int32(id)
	return pokemon, nil
}

func (repository pokemonRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Pokemon, error) {
	sqlQuery := "SELECT * FROM pokemons WHERE id = ?"
	rows, err := repository.DB.QueryContext(ctx, sqlQuery, id)
	pokemon := entity.Pokemon{}

	if err != nil {
		return pokemon, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Type, &pokemon.Species)
		return pokemon, nil
	} else {
		return pokemon, errors.New("Id :" + strconv.Itoa(int(id)) + "Not found id")
	}
}

func (repository pokemonRepositoryImpl) FindAll(ctx context.Context) ([]entity.Pokemon, error) {
	sqlQuery := "SELECT * FROM pokemons"
	rows, err := repository.DB.QueryContext(ctx, sqlQuery)
	var pokemons []entity.Pokemon

	if err != nil {
		return pokemons, err
	}
	defer rows.Close()

	for rows.Next() {
		var pokemon entity.Pokemon
		rows.Scan(&pokemon.Id, &pokemon.Name, &pokemon.Type, &pokemon.Species)
		pokemons = append(pokemons, pokemon)
	}

	return pokemons, nil
}

func (repository pokemonRepositoryImpl) UpdateById(ctx context.Context, pokemon entity.Pokemon, id int32) (entity.Pokemon, error) {
	sqlExec := "UPDATE pokemons SET name = ?, type = ?, species= ? WHERE id = ? "
	result, err := repository.DB.ExecContext(ctx, sqlExec, pokemon.Name, pokemon.Type, pokemon.Species, id)
	if err != nil {
		return pokemon, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return pokemon, errors.New("no row affected, pokemon with id " + strconv.Itoa(int(id)) + "not found")
	}

	return pokemon, nil

}

func (repository pokemonRepositoryImpl) DeleteById(ctx context.Context, id int32) error {
	sqlExec := "DELETE FROM pokemons WHERE id = ?;"
	result, err := repository.DB.ExecContext(ctx, sqlExec, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no row affected, pokemon with id " + strconv.Itoa(int(id)) + "not found")
	}

	return nil
}
