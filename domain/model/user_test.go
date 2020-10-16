package model

import "testing"

func TestNewUser(t *testing.T) {
	_, err := newUserId("")
	if err == nil {
		t.Error("error")
	}
	_, err = newEmailAddress("aaa.aaa")
	if err == nil {
		t.Error("error")
	}
	_, err = newPassword("123")
	if err == nil {
		t.Error("error")
	}
	_, err = newUserId("userId")
	if err != nil {
		t.Error(err.Error())
	}
	_, err = newEmailAddress("example@example.com")
	if err != nil {
		t.Error(err.Error())
	}
	_, err = newPassword("12345678")
	if err != nil {
		t.Error(err.Error())
	}
	_, err = NewUser("userId", "example@example.com", "12345678")
	if err != nil {
		t.Error(err.Error())
	}
}
