package main

import (
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

const defaultPort = "8080"

func main() {
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
		http.Client{}, "https://api.travelgatex.com",
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
		book.NewBookHandle(
			&book.BookService{},
		),
	)

	log.Printf("Running in port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
