package validators

import (
	"regexp"

	"github.com/dsuhinin/suhinin-backend-1/src/api"
	"github.com/dsuhinin/suhinin-backend-1/src/api/request"
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
		return api.PasswordEmptyError
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
		return api.PasswordEmptyError
	}
	confirmPassword := req.ConfirmPassword
	if confirmPassword == "" || !isPasswordValid(confirmPassword) {
		return api.ConfirmPasswordEmptyError
	}
	if req.Password != req.ConfirmPassword {
		return api.PasswordAndConfirmPasswordNotEqualError
	}
	return nil
}

// isPasswordValid check that password provided passes required length.
func isPasswordValid(password string) bool {
	if len(password) < 6 && len(password) > 254 {
		return false
	}

	return true
}

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(email string) bool {
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return emailRegex.MatchString(email)
}
