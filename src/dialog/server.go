package main

import (
	"net/http"

	"github.com/Arkiant/labxIII/src/dialog/dialogflow"

	"github.com/Arkiant/labxIII/src/dialog/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Recoverer)

	df := dialogflow.NewDialog()
	s := routes.NewServer(df)

	r.MethodFunc("POST", "/", s.RequestHandler)

	http.ListenAndServe(":6969", r)
}
