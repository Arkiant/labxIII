package main

import (
	"net/http"

	"github.com/Arkiant/labxIII/src/dialog/routes"
)

func main() {
	r := routes.NewRoutes()
	http.ListenAndServe(":5000", r)
}
