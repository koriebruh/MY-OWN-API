package repository

import (
	"api-pokemon/entity"
	"context"
)

type PokemonRepository interface {
	Insert(ctx context.Context, pokemon entity.Pokemon) (entity.Pokemon, error)
	FindById(ctx context.Context, id int32) (entity.Pokemon, error)
	FindAll(ctx context.Context) ([]entity.Pokemon, error)
	UpdateById(ctx context.Context, pokemon entity.Pokemon, id int32) (entity.Pokemon, error)
	DeleteById(ctx context.Context, id int32) error
}
