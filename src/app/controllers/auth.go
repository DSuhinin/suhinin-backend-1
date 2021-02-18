package controllers

import (
	coreResponse "github.com/dsuhinin/suhinin-backend-1/core/http/response"

	"github.com/dsuhinin/suhinin-backend-1/src/api/request"
	"github.com/dsuhinin/suhinin-backend-1/src/api/response"
)

// Signin makes processing of POST /auth/signin route.
func (c *Controller) Signin(req *request.Signin) (coreResponse.Provider, error) {

	if err := c.requestValidator.ValidateSigninRequest(req); err != nil {
		return nil, err
	}

	return coreResponse.New(response.NewSignin("token")), nil
}

// Signup makes processing of POST /auth/signup route.
func (c *Controller) Signup(req *request.Signup) (coreResponse.Provider, error) {

	if err := c.requestValidator.ValidateSignupRequest(req); err != nil {
		return nil, err
	}

	return coreResponse.New(nil), nil
}

// Signout makes processing of POST /auth/signout route.
func (c *Controller) Signout() (coreResponse.Provider, error) {
	return coreResponse.New(nil), nil
}
