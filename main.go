package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cactus/go-statsd-client/v5/statsd"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	client, err := statsd.NewClient("127.0.0.1:8125", "helloItsMe")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		client.Inc("hello.request."+name, 1, 1.0)
		fmt.Fprintf(w, "Hello, %s!", name)
	})

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", r))
}
