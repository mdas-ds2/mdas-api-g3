package main

import (
	"flag"
	"fmt"
	"log"

	application "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/application"
	pokemon "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/domain"
	infra "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/infraestructure/poke-api"
)

func main() {
	pokemonNameInput := flag.String("pokemonName", "", "The type of the pokemon we want to search for")
	flag.Parse()

	pokemonName, err := pokemon.CreatePokemonName(*pokemonNameInput)

	if err != nil {
		log.Fatalln(err)
	}

	repo := infra.PokeApiRepository{}

	getTypeByNameUseCase := application.GetTypeByName{repo}

	fmt.Println(getTypeByNameUseCase.Execute(*pokemonName))
}
