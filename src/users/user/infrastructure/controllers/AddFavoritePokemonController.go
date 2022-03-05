package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	webserver "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/web-server"
	application "github.com/mdas-ds2/mdas-api-g3/src/users/user/application"
	infrastructure "github.com/mdas-ds2/mdas-api-g3/src/users/user/infrastructure"
)

type addFavoritePokemonController struct {
	pattern string
}

const FAVORITE_POKEMON_URL_PATTERN_SEGMENT = "/favorite-pokemon/"

var InMemomyFavoritePokemonDDBB = map[string][]string{}

func (controller addFavoritePokemonController) Handler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		response.WriteHeader(http.StatusMethodNotAllowed)
		exception := webserver.CreateMethodNotSupportedException()
		webserver.RespondJsonError(response, exception.GetError())
		return
	}

	userId := request.Header.Get("UserId")
	body, err := ioutil.ReadAll(request.Body)
	fmt.Println("This is the userId ---> ", userId)
	if err != nil {
		exception := webserver.CreateInternalServerErrorException("error reading request content")
		response.WriteHeader(http.StatusInternalServerError)
		webserver.RespondJsonError(response, exception.GetError())
		return
	}

	var requestBody infrastructure.PokemonIdModel

	err = json.Unmarshal(body, &requestBody)

	if err != nil {
		exception := webserver.CreateBadRequestException("bad formatted content")
		response.WriteHeader(http.StatusBadRequest)
		webserver.RespondJsonError(response, exception.GetError())
		return
	}
	fmt.Println("This is the pokemonId ---> ", requestBody.PokemonId)
	inMemoryRepo := infrastructure.CreateFavoritePokemonMemoryRepository(&InMemomyFavoritePokemonDDBB)
	eventPublisher := application.RabbitMqEventPublisher{}
	addFavoritePokemonUseCase := application.AddFavoritePokemon{
		Repository: inMemoryRepo,
		Publisher:  eventPublisher,
	}

	error := addFavoritePokemonUseCase.Execute(userId, requestBody.PokemonId)

	if error != nil {
		response.WriteHeader(http.StatusConflict)
		webserver.RespondJsonError(response, error)
		return
	}

	respond(response, "pokemon added to favorite list correctly")
}

func (controller addFavoritePokemonController) GetPattern() string {
	return controller.pattern
}

func CreateAddFavoritePokemonController() addFavoritePokemonController {
	return addFavoritePokemonController{pattern: FAVORITE_POKEMON_URL_PATTERN_SEGMENT}
}

func respond(response http.ResponseWriter, message string) {
	resultMap := make(map[string]string)

	resultMap["response"] = message

	result, _ := json.Marshal(resultMap)

	response.Write(result)
}
