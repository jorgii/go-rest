package restapi

import "testing"

type User struct {
	Email string `validate:"email"`
}

func TestValidateStructSuccess(t *testing.T) {
	if err := ValidateStruct(&User{Email: "test@test.com"}); err != nil {
		t.Error("Expected no error from validation.")
	}
}

func TestValidateStructError(t *testing.T) {
	if err := ValidateStruct(&User{Email: "not an email"}); err == nil {
		t.Error("Expected an error from validation.")
	}
}
