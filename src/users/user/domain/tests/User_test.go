package user_test

import (
	"testing"

	domain "github.com/mdas-ds2/mdas-api-g3/src/users/user/domain"
)

func TestGetId(t *testing.T) {
	id := "1234"
	userId := domain.CreateUserId(id)
	user := domain.CreateUser(userId)

	got := user.GetId().GetValue()

	if got != id {
		t.Errorf("Did not get expected result. Wanted user ID %q, got: %q", id, got)
	}
}
