package pokeapi

import (
	"net/http"
	"strconv"

	shared "github.com/mdas-ds2/mdas-api-g3/src/shared/infrastructure"

	httpClient "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/http-client"
	domain "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/domain"
)

type PokeApiPokemonRepository struct {
	cache map[string]domain.Pokemon
}

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon/"

func CreatePokeApiPokemonRepository() PokeApiPokemonRepository {
	return PokeApiPokemonRepository{make(map[string]domain.Pokemon)}
}

func (repository PokeApiPokemonRepository) Find(id domain.Id) (domain.Pokemon, error) {
	pokemonId := strconv.Itoa(id.GetValue())

	if _, ok := repository.cache[pokemonId]; ok {
		return repository.cache[pokemonId], nil
	}

	urlPath := pokeApiUrl + pokemonId

	response, errorOnResponse := httpClient.Get(urlPath)

	if response.StatusCode == http.StatusServiceUnavailable {
		serviceUnavailableException := shared.CreatePokemonRepositoryUnavailableException()
		return domain.Pokemon{}, serviceUnavailableException.GetError()
	}

	if errorOnResponse != nil {
		return domain.Pokemon{}, errorOnResponse
	}

	if response.StatusCode == http.StatusNotFound {
		pokemonNotFoundException := domain.CreatePokemonNotFoundException(id)
		return domain.Pokemon{}, pokemonNotFoundException.GetError()
	}

	pokemon, _ := mapResponseToPokemon(response.Body)
	repository.cache[pokemonId] = pokemon

	return pokemon, nil
}

func (repository PokeApiPokemonRepository) Save(pokemon domain.Pokemon) {
	pokemonId := strconv.Itoa(pokemon.GetId().GetValue())
	repository.cache[pokemonId] = pokemon
}
