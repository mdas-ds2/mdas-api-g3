package pokemon

import (
	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
)

type IncreasePokemonAsFavorite struct {
	Repository domain.Repository
}

func (useCase IncreasePokemonAsFavorite) Execute(id int) {
	pokemonId := domain.CreateId(id)
	pokemon, _ := useCase.Repository.Find(pokemonId)
	pokemon.IncreaseFavoriteTimes()
	useCase.Repository.Save(pokemon)
}
