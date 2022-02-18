package pokemon

import (
	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
)

type GetPokemonDetails struct {
	Repository domain.Repository
}

func (getPokemonDetails *GetPokemonDetails) Execute(pokemonId int) (PokemonDetails, error) {
	id := domain.CreateId(pokemonId)
	pokemon, error := getPokemonDetails.Repository.Find(id)
	//has to return a domain exception
	if error != nil {
		return PokemonDetails{}, error
	}

	pokemonDetail := PokemonDetails{
		pokemon.GetId().GetValue(),
		pokemon.GetName().GetValue(),
		pokemon.GetHeight().GetValue(),
		pokemon.GetWeight().GetValue(),
	}

	return pokemonDetail, nil
}
