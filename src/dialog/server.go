package main

import (
	"net/http"

	"github.com/Arkiant/labxIII/src/dialog/routes"
)

func main() {
	r := routes.NewRoutes()

	/*
			err := http.ListenAndServeTLS(":443", "/etc/letsencrypt/live/enimada.com/fullchain.pem", "/etc/letsencrypt/live/enimada.com/privkey.pem", r)
		        if err != nil {
		                panic(err)
				}
	*/

	http.ListenAndServe(":6969", r)
}
