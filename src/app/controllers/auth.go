package controllers

import (
	coreResponse "github.com/dsuhinin/suhinin-backend-1/core/http/response"

	"github.com/dsuhinin/suhinin-backend-1/src/api"
	"github.com/dsuhinin/suhinin-backend-1/src/api/request"
	"github.com/dsuhinin/suhinin-backend-1/src/api/response"
	"github.com/dsuhinin/suhinin-backend-1/src/dao/auth"
)

// Signin makes processing of POST /auth/signin route.
func (c *Controller) Signin(req *request.Signin) (coreResponse.Provider, error) {

	if err := c.requestValidator.ValidateSigninRequest(req); err != nil {
		return nil, err
	}

	user, err := c.authRepository.GetByEmail(req.Email)
	if err != nil {
		return nil, api.InternalServerError.WithMessage("impossible to find user: %+v", err)
	}

	if user == nil {
		return nil, api.EntityNotFoundError("user")
	}

	if isValid := user.IsPasswordValid(req.Password); !isValid {
		return nil, api.UnauthorizedRequestError
	}

	token, err := c.jwtGenerator.Generate(user.ID, user.Email)
	if err != nil {
		return nil, api.InternalServerError.WithMessage("impossible to generate JWT token: %+v", err)
	}

	return coreResponse.New(response.NewSignin(token)), nil
}

// Signup makes processing of POST /auth/signup route.
func (c *Controller) Signup(req *request.Signup) (coreResponse.Provider, error) {

	if err := c.requestValidator.ValidateSignupRequest(req); err != nil {
		return nil, err
	}

	user, err := c.authRepository.GetByEmail(req.Email)
	if err != nil {
		return nil, api.InternalServerError.WithMessage("impossible to find user: %+v", err)
	}

	if user != nil {
		return nil, api.UserAlreadyExistsError
	}

	user, err = auth.NewUserModel(req.Email, req.Password)
	if err != nil {
		return nil, api.InternalServerError.WithMessage("impossible to create a user: %+v", err)
	}

	if err := c.authRepository.Create(user); err != nil {
		return nil, api.InternalServerError.WithMessage("impossible to create a user: %+v", err)
	}

	return coreResponse.New(nil), nil
}

// Signout makes processing of POST /auth/signout route.
func (c *Controller) Signout() (coreResponse.Provider, error) {
	return coreResponse.New(nil), nil
}
