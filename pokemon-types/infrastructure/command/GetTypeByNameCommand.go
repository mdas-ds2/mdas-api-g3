package command

import (
	"log"

	console "github.com/mdas-ds2/mdas-api-g3/generic/infrastructure/console"
	pokemonTypeUseCases "github.com/mdas-ds2/mdas-api-g3/pokemon-types/application"
	pokeApi "github.com/mdas-ds2/mdas-api-g3/pokemon-types/infrastructure/poke-api"
)

type GetTypesByNameCommand struct{}

func (command GetTypesByNameCommand) Run() {
	pokemonNameInput := console.NewInputParameter("getPokemonTypes", "Get pokemon types passing the pokemon name")
	pokeApiPokemonTypeRepository := pokeApi.PokeApiPokemonTypesRepository{}

	getByPokemonName := pokemonTypeUseCases.GetByPokemonName{
		PokemonTypeRepository: pokeApiPokemonTypeRepository,
	}

	pokemonTypes, errorOnGetPokemonTypes := getByPokemonName.Execute(pokemonNameInput)

	if errorOnGetPokemonTypes != nil {
		log.Fatalln(errorOnGetPokemonTypes)
	}

	console.Print(pokemonTypes[0].GetName())
}
