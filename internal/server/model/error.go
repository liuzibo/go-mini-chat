package model

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrUserExists      = errors.New("user already exist")
	ErrInvalidPassword = errors.New("invalid password")
)
