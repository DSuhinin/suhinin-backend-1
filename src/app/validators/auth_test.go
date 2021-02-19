// +build unit

package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dsuhinin/suhinin-backend-1/core/errors"

	"github.com/dsuhinin/suhinin-backend-1/src/api"
	"github.com/dsuhinin/suhinin-backend-1/src/api/request"
)

func TestValidator_ValidateSigninRequest_Ok(t *testing.T) {
	signinRequest := request.Signin{
		Email:    "user.email@gmail.com",
		Password: "password",
	}

	validator := NewValidator()
	assert.Nil(t, validator.ValidateSigninRequest(&signinRequest))
}

func TestValidator_ValidateSigninRequest_Error(t *testing.T) {

	type test struct {
		name    string
		error   errors.HTTPError
		request request.Signin
	}

	testList := []test{
		{
			name: "EmptyEmail",
			request: request.Signin{
				Email: "",
			},
			error: api.EmailEmptyError,
		},
		{
			name: "IncorrectEmail",
			request: request.Signin{
				Email: "incorrect.email",
			},
			error: api.EmailEmptyError,
		},
		{
			name: "EmptyPassword",
			request: request.Signin{
				Email: "user.email@gmail.com",
			},
			error: api.PasswordEmptyError,
		},
		{
			name: "SmallPassword",
			request: request.Signin{
				Email:    "user.email@gmail.com",
				Password: "p",
			},
			error: api.PasswordEmptyError,
		},
	}

	validator := NewValidator()
	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.error, validator.ValidateSigninRequest(&tt.request))
		})
	}
}

func TestValidator_ValidateSignupRequest_Ok(t *testing.T) {
	signinRequest := request.Signup{
		Email:           "user.email@gmail.com",
		Password:        "password",
		ConfirmPassword: "password",
	}

	validator := NewValidator()
	assert.Nil(t, validator.ValidateSignupRequest(&signinRequest))
}

func TestValidator_ValidateSignupRequest_Error(t *testing.T) {

	type test struct {
		name    string
		error   errors.HTTPError
		request request.Signup
	}

	testList := []test{
		{
			name: "EmptyEmail",
			request: request.Signup{
				Email: "",
			},
			error: api.EmailEmptyError,
		},
		{
			name: "IncorrectEmail",
			request: request.Signup{
				Email: "incorrect.email",
			},
			error: api.EmailEmptyError,
		},
		{
			name: "EmptyPassword",
			request: request.Signup{
				Email: "user.email@gmail.com",
			},
			error: api.PasswordEmptyError,
		},
		{
			name: "SmallPassword",
			request: request.Signup{
				Email:    "user.email@gmail.com",
				Password: "p",
			},
			error: api.PasswordEmptyError,
		},
		{
			name: "EmptyConfirmPassword",
			request: request.Signup{
				Email:    "user.email@gmail.com",
				Password: "password",
			},
			error: api.ConfirmPasswordEmptyError,
		},
		{
			name: "SmallConfirmPassword",
			request: request.Signup{
				Email:           "user.email@gmail.com",
				Password:        "password",
				ConfirmPassword: "p",
			},
			error: api.ConfirmPasswordEmptyError,
		},
		{
			name: "NotEqualPasswordAndConfirmPassword",
			request: request.Signup{
				Email:           "user.email@gmail.com",
				Password:        "password1",
				ConfirmPassword: "password2",
			},
			error: api.PasswordAndConfirmPasswordNotEqualError,
		},
	}

	validator := NewValidator()
	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.error, validator.ValidateSignupRequest(&tt.request))
		})
	}
}
