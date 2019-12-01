package pkg

import (
	"github.com/Arkiant/labxIII/src/conversation"
	"github.com/Arkiant/labxIII/src/hotelx"
)

type BookRQ struct {
	OptionID string
}

type Response struct {
	Errors []Error
}
type Error struct {
	Code        string
	Description string
}

type DestinationSearcherResponse struct {
	hotelx.Destination
}

type SearchResponse struct {
	Response
	hotelx.Search
}

type BookResponse struct {
	Response
	hotelx.Book
}

type Service interface {
	Search(input conversation.Criteria) SearchResponse
	Book(optionID string) BookResponse
}
