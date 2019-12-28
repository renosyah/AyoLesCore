package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type (
	BaseResponse struct {
		Errors []string `json:"errors,omitempty"`
	}
	Response struct {
		Status       string `json:"status"`
		BaseResponse `json:"errors"`
		Data         interface{} `json:"result"`
	}
)

var (
	MessageGeneralError = "unexpected error, please try again latter!"
	MessageUnauthorized = "Unauthorized Session"
)

// RespondError writes / respond with JSON-formatted request of given message & http status.
func RespondError(w http.ResponseWriter, message string, status int) {
	resp := Response{
		Status: http.StatusText(status),
		Data:   nil,
		BaseResponse: BaseResponse{
			Errors: []string{
				message,
			},
		},
	}

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Println(err)
		return
	}
}

func ComparePassword(ctx context.Context, hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
