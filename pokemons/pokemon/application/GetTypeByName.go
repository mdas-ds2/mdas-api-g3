package pokemon

import pokemon "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/domain"

type GetTypeByName struct {
	pokemonRepository pokemon.PokemonRepository
}

func (getTypeByName GetTypeByName) Execute(name pokemon.PokemonName) string {
	pokemon := getTypeByName.pokemonRepository.FindByName(name)
	return pokemon.GetStringFormatedTypes()
}
