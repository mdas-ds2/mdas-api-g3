package command

import (
	"log"

	console "github.com/mdas-ds2/mdas-api-g3/generic/infrastructure/console"
	pokemonTypeUseCases "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon-types/application"
	pokeApi "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon-types/infrastructure/poke-api"
	transformers "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon-types/infrastructure/transformers"
)

type getTypesByPokemonName struct{}

func (command getTypesByPokemonName) Run() {
	pokemonNameInput := console.NewInputParameter("getPokemonTypes", "Get pokemon types passing the pokemon name")
	pokeApiPokemonTypeRepository := pokeApi.PokeApiPokemonTypesRepository{}

	getByPokemonName := pokemonTypeUseCases.GetByPokemonName{
		PokemonTypeRepository: pokeApiPokemonTypeRepository,
	}

	pokemonTypes, errorOnGetPokemonTypes := getByPokemonName.Execute(pokemonNameInput)

	if errorOnGetPokemonTypes != nil {
		log.Fatalln(errorOnGetPokemonTypes)
	}

	pokemonTypesToString := transformers.PokemonTypesToString{}

	console.Print(pokemonTypesToString.Parse(pokemonTypes))
}

func NewGetTypesByPokemonName() getTypesByPokemonName {
	return getTypesByPokemonName{}
}
