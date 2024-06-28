package middleware

import (
	"net/http"
	"strings"
)

func Public(sm *http.ServeMux) func(http.HandlerFunc) http.HandlerFunc {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.URL.Path, ".") {
				sm.ServeHTTP(w, r)
				return
			} else if strings.EqualFold(r.URL.Path, "/") {
				hf(w, r)
				return
			}

			http.NotFound(w, r)
		}
	}
}
