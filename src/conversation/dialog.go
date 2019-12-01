package conversation

import (
	"io"
)

// Intent represent a conversation intent
type Intent string

// We have 2 intents search and book
const (
	SEARCH Intent = "search"
	BOOK   Intent = "book"
)

// Dialog is a interface represent a dialog API NLP
type Dialog interface {
	Convert(io.ReadCloser) (*Criteria, error)
	Send(criteria *Criteria) (io.ReadCloser, error)
	Speak(destination string, hotelName string, amount string, optionID string) ([]byte, error)
}
