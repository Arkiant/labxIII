package main

import (
	"flag"
	"github.com/Arkiant/labxIII/src/webhook/pkg"
	"github.com/Arkiant/labxIII/src/webhook/transaction"
	thttp "github.com/Arkiant/labxIII/src/webhook/transaction/http"
	"log"
	"os"

	"github.com/Arkiant/labxIII/src/webhook/book"
	"github.com/Arkiant/labxIII/src/webhook/search"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

var hXauth string

func init() {
	flag.StringVar(&hXauth, "auth", "", "the auth token to connect with hotelx")
	flag.Parse()

}

func main() {

	const defaultPort = "8080"

	if hXauth == "" {
		panic("No auth token specified")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer,
	)

	cli := thttp.NewService(
		http.Client{}, "https://api.travelgatex.com", hXauth,
	)
	service := transaction.NewService(cli)

	router.Handle("/search",
		pkg.NewRunnerHandle(
			&search.SearchFactory{
				Transactioner: service,
			},
		),
	)
	router.Handle("/book",
		pkg.NewRunnerHandle(
			&book.BookFactory{
				Transactioner: service,
			},
		),
	)

	log.Printf("Running in port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
