package main

import (
	webServer "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/web-server"
	pokemonTypesCommands "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/infrastructure/commands"
	pokemonTypesControllers "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/infrastructure/controllers"
	pokemonsControllers "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/infrastructure/controllers"
	usersControllers "github.com/mdas-ds2/mdas-api-g3/src/users/user/infrastructure/controllers"
)

func main() {
	pokemonTypesCommands.NewGetTypesByPokemonName().Run()

	server := webServer.Create()

	server.Register(pokemonTypesControllers.CreateGetTypesByPokemonName())
	server.Register(usersControllers.CreateAddFavoritePokemonController())
	server.Register(pokemonsControllers.CreateGetPokemonDetailsController())
	server.Listen(5001)
}
