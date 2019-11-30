package search

import (
	"net/http"
)

type SearchService struct {
}

func NewSearchHandle(s *SearchService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}
