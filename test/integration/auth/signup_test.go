// +build integration

package auth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/dsuhinin/suhinin-backend-1/core/errors"
	"github.com/dsuhinin/suhinin-backend-1/core/test/helpers"

	"github.com/dsuhinin/suhinin-backend-1/src/api"
	"github.com/dsuhinin/suhinin-backend-1/src/api/request"
	"github.com/dsuhinin/suhinin-backend-1/src/router/routes"
)

// TestAuthSignupEndpoint_Ok test `GET /auth/signup` endpoint.
func TestAuthSignupEndpoint_Ok(t *testing.T) {

	data, err := json.Marshal(request.Signup{
		Email:           "test.user2@gmail.com",
		Password:        "password",
		ConfirmPassword: "password",
	})
	assert.Nil(t, err)

	httpRequest, err := http.NewRequest(
		http.MethodPost,
		helpers.GenerateTestEndpoint(
			helpers.GetServiceBaseURL(),
			routes.SignupRoute,
			nil,
		),
		bytes.NewBuffer(data),
	)
	assert.Nil(t, err)

	HTTPClient := &http.Client{
		Timeout: time.Second * 2,
	}

	httpResponse, err := HTTPClient.Do(httpRequest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.Nil(t, httpResponse.Body.Close())
}

// TestAuthSignupEndpoint_Error test `GET /auth/signup` endpoint for all negative cases.
func TestAuthSignupEndpoint_Error(t *testing.T) {

	type test struct {
		name    string
		request request.Signup
		error   errors.HTTPError
	}

	testList := []test{
		{
			name:    "EmptyEmail",
			request: request.Signup{},
			error:   api.EmailEmptyError,
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
				Email: "test.user1@gmail.com",
			},
			error: api.PasswordEmptyOrIncorrectError,
		},
		{
			name: "SmallPassword",
			request: request.Signup{
				Email:    "test.user1@gmail.com",
				Password: "p",
			},
			error: api.PasswordEmptyOrIncorrectError,
		},
		{
			name: "EmptyConfirmPassword",
			request: request.Signup{
				Email:    "test.user1@gmail.com",
				Password: "password",
			},
			error: api.ConfirmPasswordEmptyOrIncorrectError,
		},
		{
			name: "SmallConfirmPassword",
			request: request.Signup{
				Email:           "test.user1@gmail.com",
				Password:        "password",
				ConfirmPassword: "p",
			},
			error: api.ConfirmPasswordEmptyOrIncorrectError,
		},
		{
			name: "NotEqualPasswordAndConfirmPassword",
			request: request.Signup{
				Email:           "test.user1@gmail.com",
				Password:        "password1",
				ConfirmPassword: "password2",
			},
			error: api.PasswordAndConfirmPasswordNotEqualError,
		},
		{
			name: "UserAlreadyExists",
			request: request.Signup{
				Email:           "test.user1@gmail.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			error: api.UserAlreadyExistsError,
		},
	}

	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.request)
			assert.Nil(t, err)

			httpRequest, err := http.NewRequest(
				http.MethodPost,
				helpers.GenerateTestEndpoint(
					helpers.GetServiceBaseURL(),
					routes.SignupRoute,
					nil,
				),
				bytes.NewBuffer(data),
			)
			assert.Nil(t, err)

			HTTPClient := &http.Client{
				Timeout: time.Second * 2,
			}

			httpResponse, err := HTTPClient.Do(httpRequest)

			assert.Nil(t, err)
			assert.Equal(t, http.StatusBadRequest, httpResponse.StatusCode)

			data, err = ioutil.ReadAll(httpResponse.Body)
			assert.Nil(t, err)
			assert.Equal(t, tt.error.Error(), string(data))
		})
	}
}
