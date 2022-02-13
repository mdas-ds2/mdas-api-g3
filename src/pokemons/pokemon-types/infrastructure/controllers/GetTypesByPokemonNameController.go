package controllers

import (
	"net/http"
	"strings"

	webserver "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/web-server"
	application "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/application"
	pokeApi "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/infrastructure/poke-api"
	transformers "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon-types/infrastructure/transformers"
)

type getTypesByPokemonName struct {
	pattern string
}

const (
	POKEMON_URL_PATH_SEGMENT_POSITION = 2
	POKEMON_TYPES_URL_PATTERN_SEGMENT = "/pokemon-types/"
)

func (controller getTypesByPokemonName) Handler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		methodNotSupportedException := webserver.CreateMethodNotSupportedException()
		response.WriteHeader(http.StatusMethodNotAllowed)
		webserver.RespondJsonError(response, methodNotSupportedException.GetError())
		return
	}

	pokemonName := getPokemonName(*request)

	pokeApiPokemonTypeRepository := pokeApi.PokeApiPokemonTypesRepository{}
	getByPokemonNameUseCase := application.GetByPokemonName{
		Repository: pokeApiPokemonTypeRepository,
	}

	pokemonTypes, errorOnGetPokemonTypes := getByPokemonNameUseCase.Execute(string(pokemonName))

	if errorOnGetPokemonTypes != nil {
		response.WriteHeader(http.StatusNotFound)
		webserver.RespondJsonError(response, errorOnGetPokemonTypes)
		return
	}

	responseBody, errorOnCreatingResponse := (transformers.PokemonTypesToJson{}).Parse(pokemonTypes)

	if errorOnCreatingResponse != nil {
		response.WriteHeader(http.StatusInternalServerError)
		webserver.RespondJsonError(response, errorOnCreatingResponse)
		return
	}

	webserver.RespondJson(response, responseBody)
}

func (controller getTypesByPokemonName) GetPattern() string {
	return controller.pattern
}

func CreateGetTypesByPokemonName() getTypesByPokemonName {
	return getTypesByPokemonName{pattern: POKEMON_TYPES_URL_PATTERN_SEGMENT}
}

func getPokemonName(request http.Request) string {
	urlPathSegments := strings.Split(request.URL.Path, "/")
	pokemonName := urlPathSegments[POKEMON_URL_PATH_SEGMENT_POSITION]
	return pokemonName
}
