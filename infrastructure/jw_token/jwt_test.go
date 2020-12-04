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

	activationToken, err := g.GenerateActivateUserToken("user_id")
	if err != nil {
		t.Error(err.Error())
	}

	c, err := NewTokenChecker(conf.App.KeyPath)
	if err != nil {
		t.Error(err.Error())
	}

	userId, err := c.CheckActivateUserToken(activationToken)
	if err != nil {
		t.Error(err.Error())
	}

	if userId != "user_id" {
		t.Error("error")
	}

	_, err = c.CheckLoginUserToken(activationToken)
	if err == nil {
		t.Error("error")
	}

	loginToken, err := g.GenerateLoginUserToken("user_id")

	userId, err = c.CheckLoginUserToken(loginToken)
	if err != nil {
		t.Error(err.Error())
	}

	if userId != "user_id" {
		t.Error("error")
	}

	_, err = c.CheckActivateUserToken(loginToken)
	if err == nil {
		t.Error("error")
	}
}
