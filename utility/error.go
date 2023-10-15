package utility

import "errors"

var (
	ErrDatabase     = errors.New("ERR_DATABASE")
	ErrUnauthorized = errors.New("ERR_UNAUTHORIZED")
	ErrBadRequest   = errors.New("ERR_BAD_REQUEST")
)
