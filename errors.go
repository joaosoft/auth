package auth

import (
	"github.com/joaosoft/errors"
)

var (
	ErrorNotFound             = errors.New(errors.LevelError, 1, "user not found")
	ErrorInvalidBodyParameter = errors.New(errors.LevelError, 2, "invalid body parameter '%s'")
)
