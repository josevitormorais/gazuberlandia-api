package middlewares

import "net/http"

func ApplicationJson(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "json")
		h.ServeHTTP(w, r)
	})
}
