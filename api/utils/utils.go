package utils

import (
	"fmt"

	"gopkg.in/go-playground/validator.v10"
)

type errorMessage struct {
	Error            string      `json:"error"`
	ErrorDescription interface{} `json:"error_description,omitempty"`
}

func formErrorMessage(err error) errorMessage {
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
