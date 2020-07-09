package handlers

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

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
