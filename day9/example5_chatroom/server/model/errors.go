package model

import "github.com/pkg/errors"

var (
	ErrUserNotExist = errors.New("user not exist")
	ErrInvalidPassed = errors.New("passwd or name not right")
	ErrInvalidParams = errors.New("invalid params")
	ErrUserExist = errors.New("user exist")
)
