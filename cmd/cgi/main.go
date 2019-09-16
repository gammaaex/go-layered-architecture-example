package main

import (
	"log"
	"net/http/cgi"

	"github.com/go-chi/chi"
	"go-layered-architecture-example/pkg/presentation/http"
)

func main() {
	r := chi.NewRouter()
	http.Register(r)

	println("Start server")

	err := cgi.Serve(r)
	if err != nil {
		log.Fatal(err)
	}
}
