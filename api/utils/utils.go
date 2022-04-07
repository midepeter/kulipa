package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

type errorMessage struct {
	Error            string      `json:"error"`
	ErrorDescription interface{} `json:"error_description,omitempty"`
}

func FormErrorMessage(err error) errorMessage {
	var (
		e           = "bad_request"
		description interface{}
	)

	switch err.(type) {
	case validator.ValidationErrors:
		errors := map[string]interface{}{}
		for _, v := range err.(validator.ValidationErrors) {
			errors[v.Field()] = fmt.Sprintf("invalid validation on tag: %s", v.Tag())
		}

		description = errors
	default:
		description = err.Error()
	}
	return errorMessage{Error: e, ErrorDescription: description}
}

func Render(w http.ResponseWriter, code int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data == nil {
		return nil
	}
	return json.NewEncoder(w).Encode(data)
}
