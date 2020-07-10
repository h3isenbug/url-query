package handlers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func GetURLParams(r *http.Request) map[string]string {
	value := r.Context().Value("params")
	if value == nil {
		return nil
	}

	return value.(map[string]string)
}

func GorillaMuxURLParamMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "params", mux.Vars(r))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func SendError(w http.ResponseWriter, statusCode int) error {
	return SendErrorWithCustomMessage(w, statusCode, http.StatusText(statusCode))
}

func SendErrorWithCustomMessage(w http.ResponseWriter, statusCode int, message string) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(&Response{
		Message: message,
	})
}
