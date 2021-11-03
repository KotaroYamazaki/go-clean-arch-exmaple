package user

import "errors"

var (
	ErrNameRequired     = errors.New("name is required")
	ErrBirthdayRequired = errors.New("birthday is required")
)
