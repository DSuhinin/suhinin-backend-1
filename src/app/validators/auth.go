package validators

import (
	"regexp"

	"github.com/dsuhinin/suhinin-backend-1/src/api"
	"github.com/dsuhinin/suhinin-backend-1/src/api/request"
)

const (
	minPasswordLength = 6
	maxPasswordLength = 255
	minEmailLength    = 3
	maxEmailLength    = 255
)

//nolint
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// ValidateSigninRequest validates request for POST /auth/signin endpoint.
func (v *Validator) ValidateSigninRequest(req *request.Signin) error {
	email := req.Email
	if email == "" || !isEmailValid(email) {
		return api.EmailEmptyError
	}
	password := req.Password
	if password == "" || !isPasswordValid(password) {
		return api.PasswordEmptyOrIncorrectError
	}
	return nil
}

// ValidateSignupRequest validates request for POST /auth/signup endpoint.
func (v *Validator) ValidateSignupRequest(req *request.Signup) error {

	email := req.Email
	if email == "" || !isEmailValid(email) {
		return api.EmailEmptyError
	}
	password := req.Password
	if password == "" || !isPasswordValid(password) {
		return api.PasswordEmptyOrIncorrectError
	}
	confirmPassword := req.ConfirmPassword
	if confirmPassword == "" || !isPasswordValid(confirmPassword) {
		return api.ConfirmPasswordEmptyOrIncorrectError
	}
	if req.Password != req.ConfirmPassword {
		return api.PasswordAndConfirmPasswordNotEqualError
	}
	return nil
}

// isPasswordValid check that password provided passes required length.
func isPasswordValid(password string) bool {
	if len(password) < minPasswordLength || len(password) > maxPasswordLength {
		return false
	}

	return true
}

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(email string) bool {
	if len(email) < minEmailLength || len(email) > maxEmailLength {
		return false
	}
	return emailRegex.MatchString(email)
}
