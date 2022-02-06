package main

import (
	webServer "github.com/mdas-ds2/mdas-api-g3/generic/infrastructure/web-server"
	pokemonTypeCommands "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon-types/infrastructure/commands"
)

func main() {
	pokemonTypeCommands.GetTypesByName().Run()

	webServer.Create().Listen(5000)
}
