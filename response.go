package web

import (
	"encoding/json"
	"net/http"
)

const (
	// HeaderContentType is the content-type header.
	HeaderContentType = "content-type"

	// HeaderContentTypeJSON is the value for the content-type header for JSON responses.
	HeaderContentTypeJSON = "application/json"
)

// APIResponse encapsulates a HTTP response as defined by https://jsonapi.org/ .
type APIResponse struct {
	Data   interface{} `json:"data,omitempty"`
	Errors []Error     `json:"errors,omitempty"`
}

// WriteData writes a response with a 200 OK response, wrapping the JSON payload in a "data" field.
func WriteData(rw http.ResponseWriter, m interface{}) error {
	return WriteDataWithCode(rw, http.StatusOK, m)
}

// WriteDataWithCode writes a response with the given HTTP response code, wrapping the JSON payload
// in a "data" field.
func WriteDataWithCode(rw http.ResponseWriter, code int, m interface{}) error {
	decorateResponse(rw, code)

	w := APIResponse{
		Data: m,
	}

	return json.NewEncoder(rw).Encode(w)
}

// WriteError writes a response, wrapping it in an "errors" field.
func WriteError(rw http.ResponseWriter, code int, err error) error {
	decorateResponse(rw, code)

	errors := []Error{
		NewError(err),
	}

	w := APIResponse{
		Errors: errors,
	}

	return json.NewEncoder(rw).Encode(w)
}

// Error compatible with the Error Object from https://jsonapi.org/format/#errors
type Error struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

// NewError retusn an error wrapped in an Error struct.
func NewError(err error) Error {
	return Error{
		Description: err.Error(),
	}
}

func decorateResponse(rw http.ResponseWriter, code int) {
	rw.Header().Add(HeaderContentType, HeaderContentTypeJSON)
	rw.WriteHeader(code)
}
