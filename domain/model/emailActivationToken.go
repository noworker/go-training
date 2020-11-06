package model

import "golang.org/x/sys/unix"

type ActivationToken string

type EmailActivationToken struct {
	ActivationToken
	UserId
	ExpiresAt unix.Time_t
}
