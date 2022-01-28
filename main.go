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
	pokemonNameInput := flag.String("getPokemonTypes", "", "Get pokemon types passing the pokemon name")
	flag.Parse()

	pokemonName, err := pokemon.CreatePokemonName(*pokemonNameInput)

	if err != nil {
		log.Fatalln(err)
	}

	pokeApiRepository := pokeApi.PokeApiPokemonRepository{}

	getTypesByName := pokemonUseCases.GetTypesByName{
		PokemonRepository: pokeApiRepository,
	}

	pokemonTypes := getTypesByName.Execute(*pokemonName)

	fmt.Println(pokemonTypes)
}
