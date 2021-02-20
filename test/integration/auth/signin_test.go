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

// TestServiceInfoEndpoint test `GET /auth/signin` endpoint.
func TestAuthSigninEndpoint_Ok(t *testing.T) {

	data, err := json.Marshal(request.Signin{
		Email:    "test.user1@gmail.com",
		Password: "password",
	})
	assert.Nil(t, err)

	httpRequest, err := http.NewRequest(
		http.MethodPost,
		helpers.GenerateTestEndpoint(
			helpers.GetServiceBaseURL(),
			routes.SigninRoute,
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

// TestAuthSigninEndpoint_Error test `GET /auth/signin` endpoint for all negative cases.
func TestAuthSigninEndpoint_Error(t *testing.T) {

	type test struct {
		name    string
		request request.Signin
		error   errors.HTTPError
	}

	testList := []test{
		{
			name:    "EmptyEmail",
			request: request.Signin{},
			error:   api.EmailEmptyError,
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
				Email: "test.user1@gmail.com",
			},
			error: api.PasswordEmptyOrIncorrectError,
		},
		{
			name: "SmallPassword",
			request: request.Signin{
				Email:    "test.user1@gmail.com",
				Password: "p",
			},
			error: api.PasswordEmptyOrIncorrectError,
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
					routes.SigninRoute,
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
