// +build integration

package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/dsuhinin/suhinin-backend-1/core/errors"
	"github.com/dsuhinin/suhinin-backend-1/core/test/helpers"

	"github.com/dsuhinin/suhinin-backend-1/src/api"
	"github.com/dsuhinin/suhinin-backend-1/src/api/request"
	"github.com/dsuhinin/suhinin-backend-1/src/api/response"
	"github.com/dsuhinin/suhinin-backend-1/src/middleware"
	"github.com/dsuhinin/suhinin-backend-1/src/router/routes"
)

// TestAuthSignoutEndpoint_Ok test `GET /auth/signout` endpoint.
func TestAuthSignoutEndpoint_Ok(t *testing.T) {

	// 1. make a Signin action.
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

	data, err = ioutil.ReadAll(httpResponse.Body)
	assert.Nil(t, err)

	signinResponse := response.Signin{}
	assert.Nil(t, json.Unmarshal(data, &signinResponse))

	// 2. Make a Signout action.
	httpRequest, err = http.NewRequest(
		http.MethodGet,
		helpers.GenerateTestEndpoint(
			helpers.GetServiceBaseURL(),
			routes.SignoutRoute,
			nil,
		),
		bytes.NewBuffer(data),
	)
	assert.Nil(t, err)

	httpRequest.Header.Add(
		middleware.ServerAuthorizationKeyHeader,
		fmt.Sprintf("%s %s", middleware.BearerAuthorization, signinResponse.Token),
	)

	httpResponse, err = HTTPClient.Do(httpRequest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.Nil(t, httpResponse.Body.Close())
}

// TestAuthSignoutEndpoint_Error test `GET /auth/signout` endpoint for all negative cases.
func TestAuthSignoutEndpoint_Error(t *testing.T) {

	type test struct {
		name  string
		error errors.HTTPError
	}

	testList := []test{
		{
			name:  "NotLoggedInUser",
			error: api.UnauthorizedRequestError,
		},
	}

	for _, tt := range testList {
		t.Run(tt.name, func(t *testing.T) {
			httpRequest, err := http.NewRequest(
				http.MethodGet,
				helpers.GenerateTestEndpoint(
					helpers.GetServiceBaseURL(),
					routes.SignoutRoute,
					nil,
				),
				nil,
			)
			assert.Nil(t, err)

			HTTPClient := &http.Client{
				Timeout: time.Second * 2,
			}

			httpResponse, err := HTTPClient.Do(httpRequest)

			assert.Nil(t, err)
			assert.Equal(t, http.StatusUnauthorized, httpResponse.StatusCode)

			data, err := ioutil.ReadAll(httpResponse.Body)
			assert.Nil(t, err)
			assert.Equal(t, tt.error.Error(), string(data))
		})
	}
}
