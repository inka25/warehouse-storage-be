package errs

import "errors"

var (
	ErrInternalServerError = errors.New("err internal server error")
	ErrEmptyBodyRequest    = errors.New("err empty body request")
	ErrAuth                = errors.New("err auth")
	ErrInvalidCountry      = errors.New("err invalid country")
	ErrInvalidUserID       = errors.New("err invalid user id")
	ErrInvalidRequest      = errors.New("err invalid request")

	ErrNoResultFound = errors.New("err no result found")
)
