package main

import (
	"log"

	pokemonUseCases "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/application"
	console "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/infrastructure/console"
	pokeApi "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/infrastructure/poke-api"
)

func main() {
	pokemonNameInput := console.NewInputParameter("getPokemonTypes", "Get pokemon types passing the pokemon name")

	pokeApiRepository := pokeApi.PokeApiPokemonRepository{}

	getTypesByName := pokemonUseCases.GetTypesByName{
		PokemonRepository: pokeApiRepository,
	}

	pokemonTypes, errorOnGetPokemonTypes := getTypesByName.Execute(pokemonNameInput)

	if errorOnGetPokemonTypes != nil {
		log.Fatalln(errorOnGetPokemonTypes)
	}

	console.Print(pokemonTypes)
}
