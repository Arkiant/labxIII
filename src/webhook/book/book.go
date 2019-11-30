package book

import (
	"net/http"
)

type BookService struct {
}

func NewBookHandle(s *BookService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}
