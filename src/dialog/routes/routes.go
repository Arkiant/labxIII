package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Route struct {
	Name    string
	Methods string
	Path    string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Dialog", "POST", "/", RequestHandler}}

func NewRoutes() *chi.Mux {
	router := chi.NewRouter()
	for _, route := range routes {

		router.Method(route.Methods, route.Path, route.Handler)

	}
	return router

}
