package model

import "golang.org/x/sys/unix"

type ActivationToken string

type EmailActivation struct {
	ActivationToken
	UserId
	ExpiresAt unix.Time_t
}
