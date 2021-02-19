package controllers

import (
	"github.com/dsuhinin/suhinin-backend-1/core/http/response"

	"github.com/dsuhinin/suhinin-backend-1/src/api/request"
	"github.com/dsuhinin/suhinin-backend-1/src/app/validators"
	"github.com/dsuhinin/suhinin-backend-1/src/dao/auth/repository"
	"github.com/dsuhinin/suhinin-backend-1/src/dep/jwt"
)

// Provider provides an interface to work with Service Controller.
type Provider interface {
	// Signin handles POST /auth/signin route.
	Signin(req *request.Signin) (response.Provider, error)
	// Signup handles POST /auth/signup route.
	Signup(req *request.Signup) (response.Provider, error)
	// Signout handles POST /auth/signout route.
	Signout(token string) (response.Provider, error)
	// GetMembers handles GET /members route.
	GetMembers() (response.Provider, error)
}

// Controller represents Service Controller.
type Controller struct {
	jwtGenerator     jwt.Provider
	userRepository   repository.UserRepositoryProvider
	tokenRepository  repository.TokenRepositoryProvider
	requestValidator validators.RequestValidator
}

// NewController returns new Service controller instance.
func NewController(
	jwtGenerator jwt.Provider,
	requestValidator validators.RequestValidator,
	userRepository repository.UserRepositoryProvider,
	tokenRepository repository.TokenRepositoryProvider,
) *Controller {
	return &Controller{
		jwtGenerator:     jwtGenerator,
		requestValidator: requestValidator,
		userRepository:   userRepository,
		tokenRepository:  tokenRepository,
	}
}
