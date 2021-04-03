package utils

import (
	"encoding/json"
	"net/http"
)

func Render(w http.ResponseWriter, code int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data == nil {
		return nil
	}
	return json.NewEncoder(w).Encode(data)
}
