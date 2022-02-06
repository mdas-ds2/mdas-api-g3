package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	pokemonTypeUseCases "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon-types/application"
	pokeApi "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon-types/infrastructure/poke-api"
)

type getTypesByPokemonName struct {
	pattern string
}

func (controller getTypesByPokemonName) Handler(response http.ResponseWriter, request *http.Request) {
	const POKEMON_URL_PATH_SEGMENT_POSITION = 2

	if request.Method == http.MethodGet {
		urlPathSegments := strings.Split(request.URL.Path, "/")
		pokemonName := urlPathSegments[POKEMON_URL_PATH_SEGMENT_POSITION]

		pokeApiPokemonTypeRepository := pokeApi.PokeApiPokemonTypesRepository{}

		getByPokemonName := pokemonTypeUseCases.GetByPokemonName{
			PokemonTypeRepository: pokeApiPokemonTypeRepository,
		}

		pokemonTypes, errorOnGetPokemonTypes := getByPokemonName.Execute(string(pokemonName))

		if errorOnGetPokemonTypes != nil {
			log.Fatalln(errorOnGetPokemonTypes)
		}

		response.Header().Set("Content-Type", "application/json")

		var res = []string{}

		for _, value := range pokemonTypes.GetValues() {
			res = append(res, value.GetName())
		}

		jsonRes, err := json.Marshal(res)

		if err != nil {
			log.Fatalln(err)
		}

		response.Write(jsonRes)
	}
}

func (controller getTypesByPokemonName) GetPattern() string {
	return controller.pattern
}

func NewGetTypesByPokemonName() getTypesByPokemonName {
	return getTypesByPokemonName{pattern: "/pokemon-types/"}
}
