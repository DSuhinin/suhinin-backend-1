package validators

import (
	"github.com/dsuhinin/suhinin-backend-1/src/api/request"
)

// RequestValidator provides validations for all requests.
type RequestValidator interface {
	// ValidateSigninRequest validates request for POST /auth/signin endpoint.
	ValidateSigninRequest(req *request.Signin) error
	// ValidateSignupRequest validates request for POST /auth/signup endpoint.
	ValidateSignupRequest(req *request.Signup) error
}

// Validator is a Service requests validator.
type Validator struct{}

// NewValidator returns a new validator instance.
func NewValidator() *Validator {
	return &Validator{}
}
