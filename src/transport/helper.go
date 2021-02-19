package transport

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/dsuhinin/suhinin-backend-1/core/http/response"

	"github.com/dsuhinin/suhinin-backend-1/src/api"
)

// unmarshal makes unmarshalling request body to a request object.
//nolint
func unmarshal(reader io.Reader, obj interface{}) response.Provider {

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return response.New(api.JSONParsingError)
	}

	if err = json.Unmarshal(body, obj); err != nil {
		if _, ok := err.(base64.CorruptInputError); ok {
			return response.New(api.JSONBase64ParsingError.WithMessage("%+v", err))
		}
		return response.New(api.JSONParsingError)
	}

	return nil
}
