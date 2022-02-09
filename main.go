package main

import (
	webServer "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/web-server"
	pokemonTypesCommands "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/infrastructure/commands"
	pokemonTypesControllers "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/infrastructure/controllers"
)

func main() {
	pokemonTypesCommands.NewGetTypesByPokemonName().Run()

	server := webServer.Create()
	server.Register(pokemonTypesControllers.NewGetTypesByPokemonName())
	server.Listen(5001)
}
