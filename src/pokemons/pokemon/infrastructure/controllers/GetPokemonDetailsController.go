package pokemon

import (
	"net/http"
	"strconv"

	webserver "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/web-server"
	application "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/application"
	pokeApi "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/infrastructure/poke-api"
	subscribers "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/infrastructure/subscribers"
	transformers "github.com/mdas-ds2/mdas-api-g3/src/pokemons/pokemon/infrastructure/transformers"
)

type getPokemonDetailsController struct {
	pattern                    string
	getByPokemonDetailsUseCase application.GetPokemonDetails
}

const POKEMON_URL_PATTERN_SEGMENT = "/pokemon/"

func (controller getPokemonDetailsController) Handler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		methodNotSupportedException := webserver.CreateMethodNotSupportedException()
		response.WriteHeader(http.StatusMethodNotAllowed)
		webserver.RespondJsonError(response, methodNotSupportedException.GetError())
		return
	}

	pokemonId, error := getPokemonId(*request)

	if error != nil {
		response.WriteHeader(http.StatusBadRequest)
		webserver.RespondJsonError(response, error)
		return
	}

	pokemon, error := controller.getByPokemonDetailsUseCase.Execute(pokemonId)

	if error != nil {
		response.WriteHeader(http.StatusNotFound)
		webserver.RespondJsonError(response, error)
		return
	}

	responseBody, error := (transformers.PokemonToJson{}).Parse(pokemon)

	if error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		webserver.RespondJsonError(response, error)
		return
	}

	webserver.RespondJson(response, responseBody)
}

func (controller getPokemonDetailsController) GetPattern() string {
	return controller.pattern
}

func CreateGetPokemonDetailsController() getPokemonDetailsController {
	pokeApiPokemonRepository := pokeApi.CreatePokeApiPokemonRepository()
	getByPokemonDetailsUseCase := application.GetPokemonDetails{
		Repository: pokeApiPokemonRepository,
	}
	controller := getPokemonDetailsController{POKEMON_URL_PATTERN_SEGMENT, getByPokemonDetailsUseCase}

	return controller
}

func (controller getPokemonDetailsController) ListenEvents() {
	(subscribers.FavoritePokemonAddedSubscriber{}).RegisterSubscriber(controller.getByPokemonDetailsUseCase)
}
func getPokemonId(request http.Request) (int, error) {
	return strconv.Atoi(request.URL.Query()["id"][0])
}
