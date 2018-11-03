package main

import (
	"net/http"
	"os"

	users "github.com/RedVentures22/campus-2018-go"
)

func main() {
	s := users.NewStore()

	h := handler{
		store: s,
	}

	r := http.NewServeMux()

	r.HandleFunc("/health", health)
	r.Handle("/v1/users", withAuth(http.HandlerFunc(h.routerUsers)))

	err := http.ListenAndServe(os.Getenv("ADDR"), withLog(r))
	if err != nil {
		panic(err)
	}
}
