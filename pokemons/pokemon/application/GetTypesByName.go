package pokemon

import pokemon "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/domain"

type GetTypesByName struct {
	PokemonRepository pokemon.PokemonRepository
}

func (getTypesByName GetTypesByName) Execute(name pokemon.PokemonName) string {
	pokemon := getTypesByName.PokemonRepository.FindByName(name)
	return pokemon.GetStringFormatedTypes()
}
