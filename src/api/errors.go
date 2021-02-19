package api

import (
	"fmt"

	"github.com/dsuhinin/suhinin-backend-1/core/errors"
)

// Core errors
var (
	// HTTP 500 errors
	InternalServerError = errors.NewHTTP500Error(
		10000,
		"internal error happened.",
	)
	EntityNotFoundError = func(entity string) errors.HTTPError {
		return errors.NewHTTP404Error(
			10001,
			fmt.Sprintf("entity `%s` not found.", entity),
		)
	}
	UnauthorizedRequestError = errors.NewHTTP401Error(
		10002,
		"unauthorized request.",
	)
	ServerAuthorizationHeaderEmptyError = errors.NewHTTP401Error(
		10003,
		"header `Authorization` not provided, incorrect or empty",
	)
	JSONParsingError = errors.NewHTTP400Error(
		30000,
		"json parsing error.",
	)
	JSONBase64ParsingError = errors.NewHTTP400Error(
		30001,
		"one of fields is not a valid base64 string.",
	)
)

// Application errors.
var (
	// Auth errors.
	EmailEmptyError = errors.NewHTTP400Error(
		20000,
		"`email` empty or is incorrect",
	)
	PasswordEmptyError = errors.NewHTTP400Error(
		20010,
		"`password` empty or is incorrect",
	)
	ConfirmPasswordEmptyError = errors.NewHTTP400Error(
		20011,
		"`confirm_password` empty or is incorrect",
	)
	PasswordAndConfirmPasswordNotEqualError = errors.NewHTTP400Error(
		20012,
		"`password` and `confirm_password` are not equal",
	)
	UserAlreadyExistsError = errors.NewHTTP400Error(
		20020,
		"user already exists with provided `email`",
	)
)
