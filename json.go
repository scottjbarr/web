package web

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ParseJSON(r *http.Request, dst interface{}) error {
	return json.NewDecoder(r.Body).Decode(dst)
}

func ParseJSONStrict(r *http.Request, dst interface{}) error {
	contentTtype := r.Header.Get("Content-Type")

	if contentTtype != "application/json" {
		return errors.New("not application/json")
	}

	if err := ParseJSON(r, dst); err != nil {
		return err
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		return err
	}

	return validator.New().Struct(dst)
}
