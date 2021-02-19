package transport

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"strconv"

	"github.com/dsuhinin/suhinin-backend-1/core/http/response"

	"github.com/dsuhinin/suhinin-backend-1/src/api"
)

// convertStringToInt converts string to integer helper.
func convertStringToInt(str string, number *int) response.Provider {
	int, err := strconv.Atoi(str)
	if err != nil {
		return response.New(api.InternalServerError.WithMessage("%+v", err))
	}

	*number = int

	return nil
}

// unmarshal makes unmarshalling request body to a request object.
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
