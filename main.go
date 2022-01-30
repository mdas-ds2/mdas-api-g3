package main

import (
	"flag"
	"fmt"
	"log"

	pokemonUseCases "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/application"
	pokemon "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/domain"
	pokeApi "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/infrastructure/poke-api"
)

func main() {
	// TODO: Separate this responsability
	pokemonNameInput := flag.String("getPokemonTypes", "", "Get pokemon types passing the pokemon name")
	flag.Parse()

	pokemonName, errorOnCreatePokemonName := pokemon.CreatePokemonName(*pokemonNameInput)

	if errorOnCreatePokemonName != nil {
		log.Fatalln(errorOnCreatePokemonName)
	}

	pokeApiRepository := pokeApi.PokeApiPokemonRepository{}

	getTypesByName := pokemonUseCases.GetTypesByName{
		PokemonRepository: pokeApiRepository,
	}

	pokemonTypes, errorOnGetPokemonTypes := getTypesByName.Execute(*pokemonName)

	if errorOnCreatePokemonName != nil {
		log.Fatalln(errorOnGetPokemonTypes)
	}

	fmt.Println(pokemonTypes)
}
