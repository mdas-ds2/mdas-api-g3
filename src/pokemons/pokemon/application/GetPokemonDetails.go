package pokemon

import (
	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
)

type GetPokemonDetails struct {
	Repository domain.Repository
}

func (getPokemonDetails *GetPokemonDetails) Execute(pokemonId int) (PokemonDetailsDTO, error) {
	id := domain.CreateId(pokemonId)
	pokemon, error := getPokemonDetails.Repository.Find(id)

	if error != nil {
		return PokemonDetailsDTO{}, error
	}

	pokemonDetail := PokemonDetailsDTO{
		pokemon.GetId().GetValue(),
		pokemon.GetName().GetValue(),
		pokemon.GetHeight().GetValue(),
		pokemon.GetWeight().GetValue(),
		pokemon.GetTimesAsFavorite().GetValue(),
	}

	return pokemonDetail, nil
}
