package jwt_lib

import (
	"go_training/config"
	"testing"
)

func TestJWT(t *testing.T) {
	conf := config.NewConfig()
	KeyGenerator(conf)
	token, err := Generator("user_id", conf)
	if err != nil {
		t.Error(err.Error())
	}

	_, err = Checker(token, conf)
	if err != nil {
		t.Error(err.Error())
	}
}
