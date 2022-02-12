package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	webserver "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/web-server"
	userUseCases "github.com/mdas-ds2/mdas-api-g3/src/users/user/application"
	infra "github.com/mdas-ds2/mdas-api-g3/src/users/user/infrastructure"
)

type addFavoritePokemonController struct {
	pattern string
}

const FAVORITE_POKEMON_URL_PATTERN_SEGMENT = "/favorite-pokemon/"

var InMemomyFavoritePokemonDDBB = map[string][]string{}

func (controller addFavoritePokemonController) Handler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		methodNotSupportedException := webserver.CreateMethodNotSupportedException()
		webserver.RespondJsonError(response, methodNotSupportedException.GetError())
		return
	}

	userId := request.Header.Get("UserId")

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	var requestBody infra.PokemonIdModel

	err = json.Unmarshal(body, &requestBody)

	if err != nil {
		panic(err)
	}

	inMemoryRepo := infra.CreateFavoritePokemonMemoryRepository(&InMemomyFavoritePokemonDDBB)
	addFavoritePokemonUseCase := userUseCases.AddFavoritePokemon{
		Repository: inMemoryRepo,
	}

	error := addFavoritePokemonUseCase.Execute(userId, requestBody.PokemonId)
	if error != nil {
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