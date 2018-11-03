package main

import (
	"net/http"
)

var allowedTokens = []string{
	"trevor",
	"jared",
}

// withAuth will check the allowed tokens list to see if the incoming
// request has a matching Authorization header
func withAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if !contains(token, allowedTokens) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// contains will search for the given string in the list of strings
func contains(needle string, haystack []string) bool {
	for _, hay := range haystack {
		if hay == needle {
			return true
		}
	}

	return false
}
