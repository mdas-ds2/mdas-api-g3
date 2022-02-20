package pokemon

import (
	httpClient "github.com/mdas-ds2/mdas-api-g3/src/generic/infrastructure/http-client"

	"testing"
)

func TestGetPokemonType(test *testing.T) {
	//Given
	url := "http://localhost:5001/pokemon/?id=25"

	//When
	httpGetType, _ := httpClient.Get(url)

	//Then
	result := string(httpGetType.Body)
	if result != "{\"Id\":25,\"Name\":\"pikachu\",\"Height\":4,\"Weight\":60}" {
		test.Errorf("Unexpected result getting the type of pikachu")
	}
}
