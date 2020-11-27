package jw_token

import (
	"go_training/config"
	"testing"
)

func TestJWT(t *testing.T) {
	conf := config.NewDummyConfig()

	g, err := NewTokenGenerator(conf.App.KeyPath)
	if err != nil {
		t.Error(err.Error())
	}

	token, err := g.GenerateActivateUserToken("user_id")
	if err != nil {
		t.Error(err.Error())
	}

	c, err := NewTokenChecker(conf.App.KeyPath)
	if err != nil {
		t.Error(err.Error())
	}

	userId, err := c.CheckActivateUserToken(token)
	if err != nil {
		t.Error(err.Error())
	}

	if userId != "user_id" {
		t.Error("error")
	}
}
