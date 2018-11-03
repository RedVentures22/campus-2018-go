package main

import (
	"net/http"
	"os"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/health", health)

	err := http.ListenAndServe(os.Getenv("ADDR"), withLog(r))
	if err != nil {
		panic(err)
	}
}
