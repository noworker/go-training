package jw_token

import (
	"go_training/config"
	"testing"
)

func TestJWT(t *testing.T) {
	conf := config.NewDummyConfig()
	Generate(conf)
	token, err := Generate("user_id", conf)
	if err != nil {
		t.Error(err.Error())
	}
	userId, err := CheckActivateUserToken(token, conf)
	if err != nil {
		t.Error(err.Error())
	}
	if userId != "user_id" {
		t.Error("error")
	}
}
