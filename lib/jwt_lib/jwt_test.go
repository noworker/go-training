package jwt_lib

import (
	"testing"
)

func TestJWT(t *testing.T) {
	KeyGenerator()
	token, err := Generator("user_id")
	if err != nil {
		t.Error(err.Error())
	}
	println("\ntoken is: ", token)

	result, err := Checker(token)
	if err != nil {
		t.Error(err.Error())
	}
	println("\nresult is: ", result)
}