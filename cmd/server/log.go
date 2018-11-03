package main

import (
	"log"
	"net/http"
	"time"
)

func withLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		dur := time.Since(start)
		log.Println(
			"level", "info",
			"msg", "incoming request",
			"uri", r.RequestURI,
			"method", r.Method,
			"dur", dur,
		)
	})
}
