package errs

import "errors"

var (
	ErrInternalServerError = errors.New("err internal server error")
	ErrEmptyBodyRequest    = errors.New("err empty body request")
	ErrAuth                = errors.New("err auth")
	ErrInvalidCountry      = errors.New("err invalid country")
	ErrInvalidUserID       = errors.New("err invalid user id")
	ErrInvalidRequestParam     = errors.New("err invalid request param")

	ErrNoResultFound = errors.New("err no result found")
	ErrInvalidWarehouseID = errors.New("err invalid warehouse id")
	ErrInvalidProductTypeID = errors.New("err invalid product type id")
	ErrInvalidBrandID = errors.New("err invalid brand id")
	ErrInvalidPageNumber = errors.New("err invalid page number")
	ErrInvalidPageLimit = errors.New("err invalid page limit")

)
