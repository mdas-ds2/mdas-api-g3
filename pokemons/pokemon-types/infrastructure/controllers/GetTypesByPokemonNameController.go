package controllers

import (
	"errors"
	"net/http"
	"strings"

	webserver "github.com/mdas-ds2/mdas-api-g3/generic/infrastructure/web-server"
	pokemonTypeUseCases "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon-types/application"
	pokeApi "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon-types/infrastructure/poke-api"
	transformers "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon-types/infrastructure/transformers"
)

type getTypesByPokemonName struct {
	pattern string
}

const POKEMON_URL_PATH_SEGMENT_POSITION = 2

func (controller getTypesByPokemonName) Handler(response http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodGet {
		webserver.RespondJsonError(response, errors.New("Method not supported"))
		return
	}

	pokemonName := getPokemonName(*request)

	pokeApiPokemonTypeRepository := pokeApi.PokeApiPokemonTypesRepository{}
	getByPokemonName := pokemonTypeUseCases.GetByPokemonName{
		PokemonTypeRepository: pokeApiPokemonTypeRepository,
	}

	pokemonTypes, errorOnGetPokemonTypes := getByPokemonName.Execute(string(pokemonName))

	if errorOnGetPokemonTypes != nil {
		webserver.RespondJsonError(response, errorOnGetPokemonTypes)
		return
	}

	responseBody, errorOnCreatingResponse := (transformers.PokemonTypesToJson{}).Parse(pokemonTypes)

	if errorOnCreatingResponse != nil {
		webserver.RespondJsonError(response, errorOnGetPokemonTypes)
		return
	}
	webserver.RespondJson(response, responseBody)
}

func (controller getTypesByPokemonName) GetPattern() string {
	return controller.pattern
}

func NewGetTypesByPokemonName() getTypesByPokemonName {
	return getTypesByPokemonName{pattern: "/pokemon-types/"}
}

func getPokemonName(request http.Request) string {
	urlPathSegments := strings.Split(request.URL.Path, "/")
	pokemonName := urlPathSegments[POKEMON_URL_PATH_SEGMENT_POSITION]
	return pokemonName
}
