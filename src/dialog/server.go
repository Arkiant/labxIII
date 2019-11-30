package main

import (
	"net/http"

	"github.com/Arkiant/labxIII/src/dialog/routes"
)

func main() {
	r := routes.NewRoutes()
	http.ListenAndServe(":6969", r)
}
