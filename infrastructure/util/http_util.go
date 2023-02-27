package util

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type ErrorMessage struct {
	message string
	code    int
}

func WriteJSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteErrorResponse(w http.ResponseWriter, status int, message string) {
	errorMessage := ErrorMessage{
		message: message,
		code:    status,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(errorMessage)
}

func GetIdFromRequest(r *http.Request) (int64, error) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return int64(id), nil
	}

	return 0, errors.New("failed to get id from url")
}
