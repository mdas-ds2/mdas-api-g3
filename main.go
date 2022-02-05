package main

import (
	pokemonTypeCommand "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon-types/infrastructure/command"
)

func main() {
	(pokemonTypeCommand.GetTypesByName{}).Run()
}
