package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Arkiant/labxIII/src/dialog/dialogflow"

	"github.com/Arkiant/labxIII/src/dialog/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	const defaultPort = "6969"

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Recoverer)

	df := dialogflow.NewDialog()
	s := routes.NewServer(df)

	r.MethodFunc("POST", "/", s.RequestHandler)

	log.Printf("Running in port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
