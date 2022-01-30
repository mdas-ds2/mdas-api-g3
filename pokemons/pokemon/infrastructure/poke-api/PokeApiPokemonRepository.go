package pokeApi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	pokemon "github.com/mdas-ds2/mdas-api-g3/pokemons/pokemon/domain"
)

type PokeApiPokemonRepository struct{}

func (pokeApiRepository PokeApiPokemonRepository) FindByName(pokemonName pokemon.PokemonName) pokemon.Pokemon {
	// TODO: Pass http as dependency (create own http service that also contains the body parsing logic)
	// TODO: Assign api url to a constant variable
	response, errorOnResponse := http.Get("https://pokeapi.co/api/v2/pokemon/" + pokemonName.GetValue())

	if errorOnResponse != nil {
		log.Fatalln(errorOnResponse)
	}

	defer response.Body.Close()

	responseBody, errorOnResponseBody := ioutil.ReadAll(response.Body)

	if errorOnResponseBody != nil {
		log.Fatalln(errorOnResponseBody)
	}

	// TODO: Separate this responsability
	var pokemonResponse PokemonModel
	json.Unmarshal(responseBody, &pokemonResponse)

	// TODO: Try to encapsulate this logic: Mapper at infra level
	var pokemonTypes = []pokemon.PokemonType{}

	for _, pokemonTypeResponse := range pokemonResponse.Types {
		pokemonType, errorOnCreatePokemonType := pokemon.CreatePokemonType(pokemonTypeResponse.Type.Name)

		if errorOnCreatePokemonType != nil {
			log.Fatalln(errorOnCreatePokemonType)
		}

		pokemonTypes = append(pokemonTypes, *pokemonType)
	}

	pokemon, errOnCreatePokemon := pokemon.CreatePokemon(pokemonName, pokemonTypes)

	if errOnCreatePokemon != nil {
		log.Fatalln(errOnCreatePokemon)
	}

	return *pokemon
}
