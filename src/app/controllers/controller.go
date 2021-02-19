package controllers

import (
	"github.com/dsuhinin/suhinin-backend-1/core/http/response"
	"github.com/dsuhinin/suhinin-backend-1/src/dep/jwt"

	"github.com/dsuhinin/suhinin-backend-1/src/api/request"
	"github.com/dsuhinin/suhinin-backend-1/src/app/validators"
	"github.com/dsuhinin/suhinin-backend-1/src/dao/auth"
)

// Provider provides an interface to work with Service Controller.
type Provider interface {
	// Signin handles POST /auth/signin route.
	Signin(req *request.Signin) (response.Provider, error)
	// Signup handles POST /auth/signup route.
	Signup(req *request.Signup) (response.Provider, error)
	// Signout handles POST /auth/signout route.
	Signout() (response.Provider, error)
	// GetMembers handles GET /members route.
	GetMembers() (response.Provider, error)
}

// Controller represents Service Controller.
type Controller struct {
	jwtGenerator     jwt.GeneratorProvider
	authRepository   auth.RepositoryProvider
	requestValidator validators.RequestValidator
}

// NewController returns new Service controller instance.
func NewController(
	jwtGenerator jwt.GeneratorProvider,
	requestValidator validators.RequestValidator,
	authRepository auth.RepositoryProvider,
) *Controller {
	return &Controller{
		jwtGenerator:     jwtGenerator,
		requestValidator: requestValidator,
		authRepository:   authRepository,
	}
}
