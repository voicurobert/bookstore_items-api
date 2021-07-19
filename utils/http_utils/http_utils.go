package http_utils

import (
	"encoding/json"
	"github.com/voicurobert/bookstore_utils-go/rest_errors"
	"net/http"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		return
	}
}

func RespondError(w http.ResponseWriter, err rest_errors.RestError) {
	RespondJson(w, err.Status, err)
}
