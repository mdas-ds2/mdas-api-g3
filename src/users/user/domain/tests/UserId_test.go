package user_test

import (
	"testing"

	user "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

func TestCreateUserId(t *testing.T) {
	id := "1234"
	userId := user.CreateUserId(id)

	got := userId.GetValue()

	if got != id {
		t.Errorf("Did not get expected result. Wanted %q, got: %q", id, got)
	}
}
