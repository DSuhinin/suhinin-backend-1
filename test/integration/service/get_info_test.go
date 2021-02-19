// +build integration

package service

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	coreHTTP "github.com/dsuhinin/suhinin-backend-1/core/http"
	"github.com/dsuhinin/suhinin-backend-1/core/test/helpers"
)

// TestServiceInfoEndpoint test `GET /service/info` endpoint.
func TestServiceInfoEndpoint(t *testing.T) {

	httpRequest, err := http.NewRequest(
		http.MethodGet,
		helpers.GenerateTestEndpoint(
			helpers.GetServiceBaseURL(),
			coreHTTP.RouteServiceInfo,
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
	assert.Equal(t, http.StatusOK, httpResponse.StatusCode)
	assert.Nil(t, httpResponse.Body.Close())
}
